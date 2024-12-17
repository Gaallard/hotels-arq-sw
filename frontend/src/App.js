import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import LoginRegister from './components/LoginRegister/LoginRegister'; 
import Home from './components/home/home';
import MisHoteles from './components/MisHoteles/MisHoteles';
import Contenedores from './components/Contenedores/Contenedores';
import MoreInfo from './components/MoreInfo/MoreInfo.jsx';
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <Routes>
       
          <Route path="/" element={<LoginRegister />} />

         
          <Route path="/home" element={<Home />} />

      
          <Route path="/mishoteles" element={<MisHoteles />} />


          <Route path="/admincontrol" element={<MisHoteles />} />

          <Route path="/moreinfo/:id" element={<MoreInfo />} />
          
          <Route path="/contenedores" element={<Contenedores />} />

        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
