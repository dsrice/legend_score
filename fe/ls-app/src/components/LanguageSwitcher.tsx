import React from 'react';
import { useTranslation } from 'react-i18next';

const LanguageSwitcher: React.FC = () => {
  const { i18n } = useTranslation();

  const changeLanguage = (lng: string) => {
    i18n.changeLanguage(lng);
  };

  return (
    <div className="flex space-x-2">
      <button 
        className={`px-2 py-1 rounded ${i18n.language === 'en' ? 'bg-blue-700' : 'bg-blue-500'}`}
        onClick={() => changeLanguage('en')}
      >
        EN
      </button>
      <button 
        className={`px-2 py-1 rounded ${i18n.language === 'ja' ? 'bg-blue-700' : 'bg-blue-500'}`}
        onClick={() => changeLanguage('ja')}
      >
        JP
      </button>
    </div>
  );
};

export default LanguageSwitcher;