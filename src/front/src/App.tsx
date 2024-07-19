// src/App.tsx

import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './HomePage';
import InscriptionPage from './InscriptionPage';
import LoginPage from './LoginPage';
import SearchPage from './SearchPage';
import UserProfile from './UserProfile';
import MapPage from './MapPage';
import StoresPage from './StoresPage';
import WatchPage from './WatchPage';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/inscription" element={<InscriptionPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/search" element={<SearchPage/>} />
        <Route path="/profile" element={<UserProfile/>} />
        <Route path="/map" element={<MapPage/>} />
        <Route path="/stores" element={<StoresPage/>} />
        <Route path ="/watch" element={<WatchPage/>} />
      </Routes>
    </Router>
  );
};

export default App;
