// src/App.tsx

import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './HomePage';
import InscriptionPage from './InscriptionPage';
import LoginPage from './LoginPage';
import SearchPage from './SearchPage';
import UserProfile from './UserProfile';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/inscription" element={<InscriptionPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/search" element={<SearchPage/>} />
        <Route path="/profile" element={<UserProfile/>} />
      </Routes>
    </Router>
  );
};

export default App;
