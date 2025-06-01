import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { apiPost } from '../services/apiClient';
import './styles/BowlingPins.css';

// Define interfaces for our data structures
interface ThrowData {
  throwCount: number;
  throwScore: number;
  pins: number[];
  isStrike: boolean;
  isSpare: boolean;
}

interface FrameData {
  frameNumber: number;
  throws: ThrowData[];
  score: number;
  isStrike: boolean;
  isSpare: boolean;
  cumulativeScore: number;
}

const GameRegistration: React.FC = () => {
  const { userId } = useParams<{ userId: string }>();
  const navigate = useNavigate();
  const { t } = useTranslation();
  const [frames, setFrames] = useState<FrameData[]>([]);
  const [currentFrame, setCurrentFrame] = useState(0);
  const [currentThrow, setCurrentThrow] = useState(0);
  const [currentPins, setCurrentPins] = useState<number[]>(Array(10).fill(0));
  const [gameDate, setGameDate] = useState<string>(new Date().toISOString().split('T')[0]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<boolean>(false);

  // Initialize frames
  useEffect(() => {
    const initialFrames: FrameData[] = Array.from({ length: 10 }, (_, i) => ({
      frameNumber: i + 1,
      throws: [],
      score: 0,
      isStrike: false,
      isSpare: false,
      cumulativeScore: 0
    }));
    setFrames(initialFrames);
  }, []);

  // Function to go back to game list
  const handleBackToGames = () => {
    navigate(`/games/${userId}`);
  };

  // Function to handle pin input
  const handlePinInput = (pins: number, pinsArray: number[] = []) => {
    if (currentFrame >= 10) return; // Game is over

    const newFrames = [...frames];
    const frame = newFrames[currentFrame];

    // If pinsArray is empty, create a default one
    const actualPinsArray = pinsArray.length > 0 
      ? pinsArray 
      : Array.from({ length: 10 }, (_, i) => i < pins ? 1 : 0);

    // Create a new throw
    const newThrow: ThrowData = {
      throwCount: currentThrow + 1,
      throwScore: pins,
      pins: actualPinsArray,
      isStrike: pins === 10 && currentThrow === 0,
      isSpare: currentThrow === 1 && pins === 10
    };

    // Add the throw to the current frame
    frame.throws.push(newThrow);

    // Update frame data
    // If it's a regular frame (not the 10th) and has 2 throws, treat the sum as 10
    if (currentFrame < 9 && frame.throws.length === 2) {
      // If it's the second throw and the total is 10, set isSpare to true
      if (currentThrow === 1 && frame.throws[0].throwScore + pins === 10) {
        newThrow.isSpare = true;
        frame.score = 10;
      }else{
        frame.score += pins - frame.throws[0].throwScore;
      }
    } else {
      frame.score = frame.throws.reduce((sum, t) => sum + t.throwScore, 0);
    }
    frame.isStrike = newThrow.isStrike || frame.isStrike;
    frame.isSpare = newThrow.isSpare || frame.isSpare;

    // Calculate scores
    console.log(3)
    calculateScores(newFrames);

    // Move to next throw or frame
    if (currentFrame === 9) {
      // Special handling for 10th frame
      if (currentThrow === 0) {
        // After first throw in 10th frame
        if (pins === 10) {
          // Strike, get two more throws
          setCurrentThrow(1);
        } else {
          // Not a strike, get one more throw
          setCurrentThrow(1);
        }
      } else if (currentThrow === 1) {
        // After second throw in 10th frame
        if (frame.isStrike || frame.isSpare) {
          // If strike or spare, get one more throw
          setCurrentThrow(2);
        } else {
          // Frame complete, game over
          setCurrentFrame(10);
        }
      } else {
        // After third throw, game over
        setCurrentFrame(10);
      }
    } else {
      // Frames 1-9
      if (pins === 10 || currentThrow === 1) {
        // Strike or second throw, move to next frame
        setCurrentFrame(currentFrame + 1);
        setCurrentThrow(0);
      } else {
        // Move to second throw in current frame
        setCurrentThrow(1);
      }
    }

    setFrames(newFrames);
  };

  // Calculate scores including strikes and spares
  const calculateScores = (frames: FrameData[]) => {
    let cumulativeScore = 0;

    for (let i = 0; i < frames.length; i++) {
      const frame = frames[i];
      let frameScore = frame.score;

      // Add bonus for strikes and spares
      if (i < 9) { // Only for frames 1-9
        if (frame.isStrike) {
          // Strike bonus: next two throws
          const nextFrame = frames[i + 1];
          if (nextFrame.throws.length > 0) {
            frameScore += nextFrame.throws[0].throwScore;

            if (nextFrame.throws.length > 1) {
              // Second throw in next frame
              frameScore += nextFrame.throws[1].throwScore;
            } else if (nextFrame.isStrike && i < 8) {
              // If next frame is also a strike, look at the first throw of the frame after that
              const frameAfterNext = frames[i + 2];
              if (frameAfterNext.throws.length > 0) {
                frameScore += frameAfterNext.throws[0].throwScore;
              }
            }
          }
        } else if (frame.isSpare) {
          // Spare bonus: next throw
          const nextFrame = frames[i + 1];
          if (nextFrame.throws.length > 0) {
            frameScore += nextFrame.throws[0].throwScore;
          }
        }
      }

      cumulativeScore += frameScore;
      frame.cumulativeScore = cumulativeScore;
    }
  };

  // Handle game submission
  const handleSubmit = async () => {
    if (currentFrame < 10) {
      setError(t('gameRegistration.error.incompleteGame'));
      return;
    }

    setIsSubmitting(true);
    setError(null);

    try {
      // Calculate total score
      const totalScore = frames[9].cumulativeScore;

      // Prepare data for API
      const gameData = {
        user_id: parseInt(userId || '0', 10),
        game_date: gameDate,
        score: totalScore,
        frames: frames.map(frame => ({
          frame_count: frame.frameNumber,
          frame_score: frame.score,
          strike_flag: frame.isStrike,
          spare_flag: frame.isSpare,
          throws: frame.throws.map(t => ({
            throw_count: t.throwCount,
            throw_score: t.throwScore,
            strike_flag: t.isStrike,
            spare_flag: t.isSpare,
            pin_1: t.pins[0],
            pin_2: t.pins[1],
            pin_3: t.pins[2],
            pin_4: t.pins[3],
            pin_5: t.pins[4],
            pin_6: t.pins[5],
            pin_7: t.pins[6],
            pin_8: t.pins[7],
            pin_9: t.pins[8],
            pin_10: t.pins[9]
          }))
        }))
      };

      // Submit to API
      const response = await apiPost('/game/register', gameData);

      if (response.result) {
        setSuccess(true);
        setTimeout(() => {
          navigate(`/games/${userId}`);
        }, 2000);
      } else {
        setError(response.message || t('gameRegistration.error.submitFailed'));
      }
    } catch (err) {
      console.error('Error submitting game:', err);
      setError(t('gameRegistration.error.submitFailed'));
    } finally {
      setIsSubmitting(false);
    }
  };

  // Reset the game
  const handleReset = () => {
    const initialFrames: FrameData[] = Array.from({ length: 10 }, (_, i) => ({
      frameNumber: i + 1,
      throws: [],
      score: 0,
      isStrike: false,
      isSpare: false,
      cumulativeScore: 0
    }));
    setFrames(initialFrames);
    setCurrentFrame(0);
    setCurrentThrow(0);
    setError(null);
    setSuccess(false);
  };

  // Render pin display
  const renderPinDisplay = () => {
    if (currentFrame >= 10) return null;

    const frame = frames[currentFrame];
    const maxPins = 10;
    const pinsRemaining = currentThrow === 0 ? maxPins : 
                          (currentFrame === 9 && (frame.isStrike || frame.isSpare)) ? maxPins : 
                          maxPins - (frame.throws[0]?.throwScore || 0);

    // Get pins that were knocked down in the first throw (for second throw)
    // In the backend, 1 means knocked down, 0 means standing
    const firstThrowKnockedDownPins = currentThrow === 1 && frame.throws.length > 0
      ? frame.throws[0].pins
      : Array(10).fill(0);

    // Function to handle individual pin clicks
    const handlePinClick = (pinIndex: number) => {
      // If it's the second throw and this pin was already knocked down in the first throw, do nothing
      if (currentThrow === 1 && firstThrowKnockedDownPins[pinIndex] === 1) {
        return;
      }

      const newPins = [...currentPins];
      newPins[pinIndex] = newPins[pinIndex] === 0 ? 1 : 0;

      // Count how many pins are remaining (pins with value 1)
      const remainingPinsCount = newPins.filter(pin => pin === 1).length;

      // Validate that we don't have more remaining pins than allowed
      if (remainingPinsCount > pinsRemaining) {
        return; // Can't have more remaining pins than allowed
      }

      setCurrentPins(newPins);
    };

    // Function to submit the current pin selection
    const submitPinSelection = () => {
      // Count how many pins are remaining (pins with value 1)
      const remainingPinsCount = currentPins.filter(pin => pin === 1).length;
      // Calculate knocked down pins as 10 minus remaining pins
      const knockedDownCount = 10 - remainingPinsCount;

      // We also need to invert the pins array for the backend
      // In the backend, 1 means knocked down, 0 means standing
      const invertedPinsArray = currentPins.map(pin => pin === 1 ? 0 : 1);

      handlePinInput(knockedDownCount, invertedPinsArray);
      // Reset pins for next throw
      setCurrentPins(Array(10).fill(0));
    };

    return (
      <div className="mt-6">
        <h3 className="text-lg font-medium text-gray-900 mb-2">
          {t('gameRegistration.selectPins')}
        </h3>

        {/* Bowling pin display */}
        <div className="bowling-pins-container mb-4">
          {/* Row 1 (back row - pins 7,8,9,10) */}
          <div className="bowling-pins-row">
            {[6, 7, 8, 9].map((pinIndex) => (
              <div 
                key={pinIndex}
                onClick={() => handlePinClick(pinIndex)}
                className={`bowling-pin ${
                  // If it's the second throw and this pin was already knocked down in the first throw, disable it
                  // but only show as knocked down if it was knocked down in the current throw
                  currentThrow === 1 && firstThrowKnockedDownPins[pinIndex] === 1
                    ? 'disabled'
                    : currentPins[pinIndex] === 1 
                      ? 'knocked-down' 
                      : 'standing'
                }`}
              >
                {pinIndex + 1}
              </div>
            ))}
          </div>

          {/* Row 2 (pins 4,5,6) */}
          <div className="bowling-pins-row">
            {[3, 4, 5].map((pinIndex) => (
              <div 
                key={pinIndex}
                onClick={() => handlePinClick(pinIndex)}
                className={`bowling-pin ${
                  currentThrow === 1 && firstThrowKnockedDownPins[pinIndex] === 1
                    ? 'disabled'
                    : currentPins[pinIndex] === 1 
                      ? 'knocked-down' 
                      : 'standing'
                }`}
              >
                {pinIndex + 1}
              </div>
            ))}
          </div>

          {/* Row 3 (pins 2,3) */}
          <div className="bowling-pins-row">
            {[1, 2].map((pinIndex) => (
              <div 
                key={pinIndex}
                onClick={() => handlePinClick(pinIndex)}
                className={`bowling-pin ${
                  currentThrow === 1 && firstThrowKnockedDownPins[pinIndex] === 1
                    ? 'disabled'
                    : currentPins[pinIndex] === 1 
                      ? 'knocked-down' 
                      : 'standing'
                }`}
              >
                {pinIndex + 1}
              </div>
            ))}
          </div>

          {/* Row 4 (pin 1) */}
          <div className="bowling-pins-row">
            <div 
              onClick={() => handlePinClick(0)}
              className={`bowling-pin ${
                currentThrow === 1 && firstThrowKnockedDownPins[0] === 1
                  ? 'disabled'
                  : currentPins[0] === 1 
                    ? 'knocked-down' 
                    : 'standing'
              }`}
            >
              1
            </div>
          </div>
        </div>

        {/* Submit button */}
        <button
          onClick={submitPinSelection}
          className="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          {t('gameRegistration.submitPins')}
        </button>

        {/* Strike button - only show on first throw */}
        {currentThrow === 0 && (
          <div className="mt-4">
            <button
              onClick={() => {
                // Set all pins to 0 (all pins knocked down, none remaining)
                setCurrentPins(Array(10).fill(0));
              }}
              className="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
            >
              {t('gameRegistration.strikeButton')}
            </button>
          </div>
        )}
      </div>
    );
  };

  // Function to check if a frame's score is confirmed
  const isFrameScoreConfirmed = (frameIndex) => {
    const frame = frames[frameIndex];

    // If no throws in this frame, score is not confirmed
    if (frame.throws.length === 0) return false;

    // For the 10th frame
    if (frameIndex === 9) {
      // 10th frame needs all throws completed
      if (frame.isStrike || frame.isSpare) {
        return frame.throws.length === 3;
      } else {
        return frame.throws.length === 2;
      }
    }

    // For frames 1-9
    if (frame.isStrike) {
      // Strike needs next two throws
      let nextThrowsCount = 0;

      // Check next frame
      if (frameIndex < 9 && frames[frameIndex + 1].throws.length > 0) {
        nextThrowsCount++;

        if (frames[frameIndex + 1].isStrike) {
          // If next frame is a strike, we need one throw from the frame after that
          if (frameIndex < 8 && frames[frameIndex + 2].throws.length > 0) {
            nextThrowsCount++;
          }
        } else if (frames[frameIndex + 1].throws.length > 1) {
          // If next frame is not a strike, we need the second throw from that frame
          nextThrowsCount++;
        }
      }

      return nextThrowsCount === 2;
    } else if (frame.isSpare) {
      // Spare needs one throw from next frame
      return frameIndex < 9 && frames[frameIndex + 1].throws.length > 0;
    } else {
      // Regular frame needs both throws
      return frame.throws.length === 2;
    }
  };

  // Render the scorecard
  const renderScorecard = () => {
    return (
      <div className="mt-6 overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              {frames.map((frame) => (
                <th key={frame.frameNumber} className="px-3 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider border">
                  {t('gameRegistration.frame')} {frame.frameNumber}
                </th>
              ))}
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            <tr>
              {frames.map((frame) => (
                <td key={`throws-${frame.frameNumber}`} className="px-3 py-4 text-center text-sm text-gray-900 border">
                  <div className="grid grid-cols-2 gap-1">
                    {frame.throws.map((t, idx) => (
                      <div key={idx} className={`p-1 ${t.isStrike ? 'bg-green-100' : t.isSpare ? 'bg-blue-100' : ''}`}>
                        {t.isStrike ? 'X':
                         (t.isSpare ? '/' : (idx === 0 ? t.throwScore: t.pins[1]))}
                      </div>
                    ))}
                    {frame.throws.length === 0 && <div className="p-1">-</div>}
                    {frame.throws.length === 1 && frame.frameNumber !== 10 && !frame.isStrike && <div className="p-1">-</div>}
                  </div>
                </td>
              ))}
            </tr>
            <tr>
              {frames.map((frame, index) => (
                <td key={`score-${frame.frameNumber}`} className="px-3 py-4 text-center text-sm font-bold text-gray-900 border">
                  {isFrameScoreConfirmed(index) ? frame.cumulativeScore : ''}
                </td>
              ))}
            </tr>
          </tbody>
        </table>
      </div>
    );
  };

  return (
    <div className="py-8 px-4 sm:px-6 lg:px-8">
      <div className="max-w-7xl mx-auto">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-2xl font-bold text-gray-900">
            {t('gameRegistration.title')}
          </h1>
          <button
            onClick={handleBackToGames}
            className="py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            {t('gameRegistration.backToGames')}
          </button>
        </div>

        {/* Error message */}
        {error && (
          <div className="bg-red-50 border-l-4 border-red-400 p-4 mb-6">
            <div className="flex">
              <div className="flex-shrink-0">
                <svg className="h-5 w-5 text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                  <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                </svg>
              </div>
              <div className="ml-3">
                <p className="text-sm text-red-700">{error}</p>
              </div>
            </div>
          </div>
        )}

        {/* Success message */}
        {success && (
          <div className="bg-green-50 border-l-4 border-green-400 p-4 mb-6">
            <div className="flex">
              <div className="flex-shrink-0">
                <svg className="h-5 w-5 text-green-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                  <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
                </svg>
              </div>
              <div className="ml-3">
                <p className="text-sm text-green-700">{t('gameRegistration.success')}</p>
              </div>
            </div>
          </div>
        )}

        <div className="bg-white shadow overflow-hidden rounded-lg p-6">
          {/* Game date input */}
          <div className="mb-6">
            <label htmlFor="gameDate" className="block text-sm font-medium text-gray-700">
              {t('gameRegistration.gameDate')}
            </label>
            <input
              type="date"
              id="gameDate"
              name="gameDate"
              value={gameDate}
              onChange={(e) => setGameDate(e.target.value)}
              className="mt-1 block w-full py-2 px-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
            />
          </div>

          {/* Current frame and throw info */}
          {currentFrame < 10 && (
            <div className="mb-6">
              <h2 className="text-xl font-semibold text-gray-900">
                {t('gameRegistration.currentFrame')}: {currentFrame + 1}, {t('gameRegistration.currentThrow')}: {currentThrow + 1}
              </h2>
            </div>
          )}

          {/* Scorecard */}
          {renderScorecard()}

          {/* Pin selection */}
          {renderPinDisplay()}

          {/* Action buttons */}
          <div className="mt-8 flex justify-between">
            <button
              onClick={handleReset}
              className="py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              disabled={isSubmitting}
            >
              {t('gameRegistration.reset')}
            </button>
            <button
              onClick={handleSubmit}
              className={`py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white ${
                currentFrame < 10 ? 'bg-gray-400 cursor-not-allowed' : 'bg-blue-600 hover:bg-blue-700'
              } focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500`}
              disabled={currentFrame < 10 || isSubmitting}
            >
              {isSubmitting ? t('gameRegistration.submitting') : t('gameRegistration.register')}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default GameRegistration;