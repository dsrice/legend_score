import React from 'react';
import { useNavigate } from 'react-router-dom';
import { removeToken } from '../services/auth';
import { useTranslation } from 'react-i18next'; // Import useTranslation hook
import LanguageSwitcher from './LanguageSwitcher';

interface HeaderProps {
  pageName: string;
}

const Header: React.FC<HeaderProps> = ({ pageName }) => {
  const navigate = useNavigate();
  const { t } = useTranslation(); // Initialize translation function

  // Translate the page name based on the current route
  const getTranslatedPageName = () => {
    if (pageName === 'Home') return t('header.home');
    if (pageName === 'User Management') return t('header.userManagement');
    return pageName;
  };

  return (
    <header className="bg-blue-600 text-white shadow-md">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div className="flex justify-between items-center">
          <h1 className="text-xl font-bold">{getTranslatedPageName()}</h1>
          <LanguageSwitcher />
        </div>
      </div>
    </header>
  );
};

export default Header;