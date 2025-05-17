import React from 'react';
import { Navigate } from 'react-router-dom';
import { isAuthenticated } from '../services/auth';
import AuthenticatedLayout from '../components/AuthenticatedLayout';

interface PrivateRouteProps {
  children: React.ReactNode;
}

/**
 * A wrapper component that redirects to the login page if the user is not authenticated
 */
const PrivateRoute: React.FC<PrivateRouteProps> = ({ children }) => {
  const authenticated = isAuthenticated();

  if (!authenticated) {
    return <Navigate to="/" replace />;
  }

  return <AuthenticatedLayout>{children}</AuthenticatedLayout>;
};

export default PrivateRoute;