import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

const GameRegistration: React.FC = () => {
  const { userId } = useParams<{ userId: string }>();
  const navigate = useNavigate();
  const { t } = useTranslation();

  // Function to go back to game list
  const handleBackToGames = () => {
    navigate(`/games/${userId}`);
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

        {/* ゲーム登録ボタンのみを表示 */}
        <div className="bg-white shadow overflow-hidden rounded-lg p-6">
          <div className="flex justify-center">
            <button
              className="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              {t('gameRegistration.register')}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default GameRegistration;