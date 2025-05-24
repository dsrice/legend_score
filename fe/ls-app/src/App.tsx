import React, { useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, useNavigate } from 'react-router-dom';
import './App.scss';

// Pages
import Login from './pages/Login';
import Home from './pages/Home';
import UserList from './pages/UserList';
import GameList from './pages/GameList';

// Utils
import PrivateRoute from './utils/PrivateRoute';
import { registerNavigationCallback } from './utils/navigation';

// NavigationProvider component to register the navigation callback
const NavigationProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const navigate = useNavigate();

  useEffect(() => {
    registerNavigationCallback((path) => {
      navigate(path);
    });
  }, [navigate]);

  return <>{children}</>;
};

function App() {
  return (
    <Router>
      <NavigationProvider>
        <div className="App">
          <Routes>
            <Route path="/" element={<Login />} />
            <Route 
              path="/home" 
              element={
                <PrivateRoute>
                  <Home />
                </PrivateRoute>
              } 
            />
            <Route 
              path="/users" 
              element={
                <PrivateRoute>
                  <UserList />
                </PrivateRoute>
              } 
            />
            {/* Add new route for GameList */}
            <Route 
              path="/games/:userId" 
              element={
                <PrivateRoute>
                  <GameList />
                </PrivateRoute>
              } 
            />
          </Routes>
        </div>
      </NavigationProvider>
    </Router>
  );
}

export default App;