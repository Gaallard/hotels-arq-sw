import React, { useState } from "react";
import swal from 'sweetalert2';
import './LoginRegister.css';
import { FaUserAlt, FaLock } from "react-icons/fa";
import axios from 'axios';

const LoginRegister = () => {
  const [action, setAction] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [registerUsername, setRegisterUsername] = useState('');
  const [registerPassword, setRegisterPassword] = useState('');

  const registerLink = () => {
    setAction('active');
  };

  const loginLink = () => {
    setAction('');
  };

  /*
  const handleRegister = (e) => {
    e.preventDefault();
 
    if (registerUsername && registerPassword) {
      swal.fire("Registro exitoso", "Bienvenido", "success").then(() => {
        window.location.href = "/users";
      });
    } else {
      swal.fire("Error", "No se pudo registrar al usuario", "error");
    }
  };*/

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/users/login', {
        username,
        password
      });
  
      if (response.status === 200) {
        const { idUser, role, token } = response.data;
        sessionStorage.setItem('token', token);
        sessionStorage.setItem('role', role);
  
        swal.fire("Login exitoso", "Bienvenido", "success").then(() => {
          if (role === 1) { // Asumiendo que role = 1 es admin
            window.location.href = "/AdminControl";
          } else {
            window.location.href = "/courses";
          }
        });
  
      } else {
        swal.fire("Error", "Usuario o contraseña incorrectos", "error");
      }
    } catch (error) {
      console.error("Error during login:", error);
      swal.fire('Error durante el login', error.message, 'error');
    }
  };
  
  
  const handleRegister = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/users/register', {
        username: registerUsername,
        password: registerPassword
      });
      if (response.status === 201) {
        swal.fire("Registro exitoso", "Bienvenido", "success").then(() => {
          window.location.href = "/users";
        });
      } else if (response.status === 400) {
        swal.fire("Error", "El usuario ya existe", "error");
      } else {
        swal.fire("Error", "No se pudo registrar al usuario", "error");
      }
    } catch (error) {
      console.error("Error during registration:", error);
      alert('Error during registration');
    }
  };
  

  const isLoginComplete = username && password;
  const isRegisterComplete = registerUsername && registerPassword;

  return (
    <div className={`wrapper ${action}`}>
      <div className="from-box login">
        <form onSubmit={handleLogin}>
          <h1>Login</h1>
          <div className="input-box">
            <input 
              type="text" 
              placeholder="Usuario" 
              required 
              value={username} 
              onChange={(e) => setUsername(e.target.value)} 
            />
            <FaUserAlt className="icon" />
          </div>
          <div className="input-box">
            <input 
              type="password" 
              placeholder="Contraseña" 
              required 
              value={password} 
              onChange={(e) => setPassword(e.target.value)} 
            />
            <FaLock className="icon" />
          </div>
          <button type="submit" disabled={!isLoginComplete}>Login</button>
          <div className="register">
            <p>No tienes una cuenta? <a href="#" onClick={registerLink}>Regístrate</a></p>
          </div>
        </form>
      </div>

      <div className="from-box register">
        <form onSubmit={handleRegister}>
          <h1>Regístrate</h1>
          <div className="input-box">
            <input 
              type="text" 
              placeholder="Usuario" 
              required 
              value={registerUsername} 
              onChange={(e) => setRegisterUsername(e.target.value)} 
            />
            <FaUserAlt className="icon" />
          </div>
          <div className="input-box">
            <input 
              type="password" 
              placeholder="Contraseña" 
              required 
              value={registerPassword} 
              onChange={(e) => setRegisterPassword(e.target.value)} 
            />
            <FaLock className="icon" />
          </div>
          <button type="submit" disabled={!isRegisterComplete}>Registrarse</button>
          <div className="register-link">
            <p>¿Ya tienes una cuenta? <a href="#" onClick={loginLink}>Login</a></p>
          </div>
        </form>
      </div>
    </div>
  );
};

export default LoginRegister;
