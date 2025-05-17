import React from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { removeToken } from '../services/auth';
import { useTranslation } from 'react-i18next'; // Import useTranslation hook

const Sidebar: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const { t } = useTranslation(); // Initialize translation function

  // Check if the current path matches the given path
  const isActive = (path: string) => {
    return location.pathname === path;
  };

  const handleLogout = () => {
    removeToken();
    navigate('/');
  };

  return (
    <div className="bg-gray-800 text-white w-64 min-h-screen flex-shrink-0 flex flex-col">
      <div className="p-4">
        <h2 className="text-xl font-semibold">{t('appName')}</h2>
      </div>
      <nav className="mt-5 flex-grow">
        <ul>
          <li>
            <button
              onClick={() => navigate('/home')}
              className={`flex items-center w-full px-4 py-3 text-left ${
                isActive('/home') 
                  ? 'bg-blue-600' 
                  : 'hover:bg-gray-700'
              }`}
            >
              <svg 
                xmlns="http://www.w3.org/2000/svg" 
                className="h-5 w-5 mr-3" 
                fill="none" 
                viewBox="0 0 24 24" 
                stroke="currentColor"
              >
                <path 
                  strokeLinecap="round" 
                  strokeLinejoin="round" 
                  strokeWidth={2} 
                  d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" 
                />
              </svg>
              {t('navigation.home')}
            </button>
          </li>
          <li>
            <button
              onClick={() => navigate('/users')}
              className={`flex items-center w-full px-4 py-3 text-left ${
                isActive('/users') 
                  ? 'bg-blue-600' 
                  : 'hover:bg-gray-700'
              }`}
            >
              <svg 
                xmlns="http://www.w3.org/2000/svg" 
                className="h-5 w-5 mr-3" 
                fill="none" 
                viewBox="0 0 24 24" 
                stroke="currentColor"
              >
                <path 
                  strokeLinecap="round" 
                  strokeLinejoin="round" 
                  strokeWidth={2} 
                  d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" 
                />
              </svg>
              {t('navigation.userManagement')}
            </button>
          </li>
        </ul>
      </nav>

      {/* Logout button at the bottom of sidebar */}
      <div className="p-4 border-t border-gray-700">
        <button
          onClick={handleLogout}
          className="flex items-center w-full px-4 py-2 text-left text-red-300 hover:bg-gray-700 rounded"
        >
          <svg 
            xmlns="http://www.w3.org/2000/svg" 
            className="h-5 w-5 mr-3" 
            fill="none" 
            viewBox="0 0 24 24" 
            stroke="currentColor"
          >
            <path 
              strokeLinecap="round" 
              strokeLinejoin="round" 
              strokeWidth={2} 
              d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" 
            />
          </svg>
          {t('navigation.logout')}
        </button>
      </div>
    </div>
  );
};

export default Sidebar;