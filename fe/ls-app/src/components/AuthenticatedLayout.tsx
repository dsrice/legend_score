import React from 'react';
import Header from './Header';
import Sidebar from './Sidebar';
import { useLocation } from 'react-router-dom';

interface AuthenticatedLayoutProps {
  children: React.ReactNode;
}

const AuthenticatedLayout: React.FC<AuthenticatedLayoutProps> = ({ children }) => {
  const location = useLocation();

  // Map routes to page names
  const getPageName = () => {
    switch (location.pathname) {
      case '/home':
        return 'Home';
      case '/users':
        return 'User List';
      default:
        return 'Legend Score';
    }
  };

  return (
    <div className="min-h-screen flex flex-col">
      {/* Header Area */}
      <div className="header-area">
        <Header pageName={getPageName()} />
      </div>

      {/* Screen Area - with horizontal layout */}
      <div className="screen-area flex flex-grow">
        {/* Side Menu */}
        <div className="side-menu-container">
          <Sidebar />
        </div>

        {/* Main Content */}
        <main className="content-container flex-grow bg-gray-50 overflow-auto p-4">
          {children}
        </main>
      </div>
    </div>
  );
};

export default AuthenticatedLayout;