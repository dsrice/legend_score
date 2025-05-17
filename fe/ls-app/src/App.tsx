import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.scss';

// Pages
import Login from './pages/Login';
import Home from './pages/Home';
import UserList from './pages/UserList';

// Utils
import PrivateRoute from './utils/PrivateRoute';

function App() {
  return (
    <Router>
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
        </Routes>
      </div>
    </Router>
  );
}

export default App;