import React from 'react';

const Home: React.FC = () => {
  return (
    <div className="py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md mx-auto space-y-8">
        <div>
          <h2 className="text-center text-3xl font-extrabold text-gray-900">
            Welcome to Legend Score
          </h2>
          <p className="mt-2 text-center text-sm text-gray-600">
            You are successfully logged in!
          </p>
        </div>
        {/* Rest of the content */}
      </div>
    </div>
  );
};

export default Home;