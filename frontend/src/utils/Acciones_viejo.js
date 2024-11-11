import axios from 'axios';
const authToken = localStorage.getItem('token');
//axios.defaults.withCredentials = true;
 
/*
export function login(userData) {
  return axios.post('http://localhost:8080/users/login', userData,{
    credentials: "include",
  })
  .then(response => {
    console.log('Login response: ', response)
    localStorage.setItem('token',response.data.Token)
    return response.data.Token;
  })
  .catch(error => {
    console.error('Login error: ', error)
    throw error;
  });
} */

export function login(loginRequest){
    return axios.post('http://localhost:8081/users/login', loginRequest) // Added http://
    .then(function (loginResponse) {
        console.log("Token: ", loginResponse.data);
        return loginResponse.data;
    })
    .catch(function(error){
        console.log("Error en el logueo", error);
        throw error;
    });
}

/*
export function register(userData){
  return axios.post('http://localhost:8080/users', userData)
    .then(response => {
      console.log('Register response:', response)
      return response.data
    })
    .catch(error => {
      console.error('Register error:', error)
      throw error;
    });
}*/

export function registration(userRequest){
    return axios
    .post('http://localhost:8081/users/register', userRequest)
    .then(function (userResponse) {
        console.log("Token: ", userResponse.data);
        return userResponse.data;
    })
    .catch(function(error){
        console.log("Error en la registracion", error);
        throw error;
    });
}

/*
export function insertHotel(Data){
  return axios.post('http://localhost:8080/courses',Data, {
    headers: { 'Authorization': `Bearer ${authToken}` }
  })
  .then(response => {
    console.log('Hotel inserted succes: ', response)
    return response.data
  })
  .catch(error => {
    console.error('Hotel error: ', error)
    throw error;
  });
}*/


export function insertHotel(hotelRequest){
    return axios
    .post('http://localhost:8080/hotels', hotelRequest)
    .then(function (hotelResponse) {
        console.log("Token: ", hotelResponse.data);
        return hotelResponse.data;
    })
    .catch(function(error){
        console.log("Error creating hotel", error);
        throw error;
    });
}

/*
export function updateHotel(Data){
  return axios.update(`http://localhost:8080/hotels/update/${Data}`, {
    headers: { 'Authorization': `Bearer ${authToken}` }
  })
  .then(response => {
    console.log('Updated hotel: ', response)
    return response.data
  })
  .catch(error => {
    console.error('hotel error: ', error)
    throw error;
  });
}*/

export function updateHotel(hotelRequest){
    return axios
    .put('http://localhost:8080/courses/update', hotelRequest)
    .then(function (hotelResponse) {
        console.log("Token: ", hotelResponse.data);
        return hotelResponse.data;
    })
    .catch(function(error){
        console.log("Error en la creacion de cursos", error);
        throw error;
    });
}

export function insertReserva(userRequest){
    return axios
    .post('http://localhost:8082/reservas', userRequest)
    .then(function (userResponse) {
        console.log("Token: ", userResponse.data);
        return userResponse.data;
    })
    .catch(function(error){
        console.log("Error en la registracion", error);
        throw error;
    });
}

export function updateReserva(userRequest){
    return axios
    .put('http://localhost:8082/reservas', userRequest)
    .then(function (userResponse) {
        console.log("Token: ", userResponse.data);
        return userResponse.data;
    })
    .catch(function(error){
        console.log("Error actualizando la reserva", error);
        throw error;
    });
}


export function getHotels() {
    const token = sessionStorage.getItem('token');
  
    return fetch('http://localhost:8080/hotels', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`
      }
    })
    .then(response => {
      if (!response.ok) {
        throw new Error('Error en la carga de hoteles');
      }
      return response.json();
    })
    .then(data => {
      return data.results;
    })
    .catch(error => {
      console.log('Error en la carga de hoteles', error);
      throw error;
    });
  }

  /*
export function getHotelsByUser(){
    return axios
    .get('http://localhost:8080/hotels/:idUser')
    .then(function (response) {
        return response.data.results;
    })
    .catch(function(error){
        console.log("Error en la carga de hoteles", error);
        throw error;
    });
}*/