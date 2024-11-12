import axios from 'axios';
const authToken = localStorage.getItem('token');
//axios.defaults.withCredentials = true;
 
export async function login(userData) {
  try {
    const response = await axios.post('http://localhost:8080/users/login', userData, {
      credentials: "include",
    });
    console.log('Login response: ', response);
    localStorage.setItem('token', response.data.Token);
    return response.data.Token;
  } catch (error) {
    console.error('Login error: ', error);
    throw error;
  }
}
 
export async function register(userData){
  try {
    const response = await axios.post('http://localhost:8080/users', userData);
    console.log('Register response:', response);
    return response.data;
  } catch (error) {
    console.error('Register error:', error);
    throw error;
  }
}

export async function insertHotel({ name, address, country, city, state, amenities, rating, price, available_rooms }) {
  try {
      console.log("Enviando datos al servidor:", { name, address, country, city, state, amenities, rating, price, available_rooms });
      const response = await axios.post('http://localhost:8081/hotels', 
          { name, address, country, city, state, amenities, rating, price, available_rooms }, 
          {
              headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
          });
      console.log("Respuesta del servidor:", response);
      return response.data;
  } catch (error) {
      console.error('Error al crear hotel en Acciones.js:', error);
      throw error;
  }
}


export async function updateHotel(hotelId, { name, address, country, city, state, amenities, rating, price, available_rooms }) {
  try {
    const response = await axios.put(`http://localhost:8081/hotels/${hotelId}`, { name, address, country, city, state, amenities, rating, price, available_rooms }, {
      headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
    });
    return response.data;
  } catch (error) {
    console.error('Error al actualizar el hotel:', error);
    throw error;
  }
}


// Obtiene todos los hoteles
export async function getAllHotels() {
  try {
    const response = await axios.get('http://localhost:8081/hotels', {
      headers: { 'Authorization': `Bearer ${authToken}` }
    });
    console.log('Hoteles cargados:', response.data);
    return response.data;
  } catch (error) {
    console.error('Error al obtener los hoteles:', error.response ? error.response.data : error.message);
    throw error;
  }
}

export async function getHotelById(hotelId) {
  try {
    const response = await axios.get(`http://localhost:8081/hotels/${hotelId}`, {
      headers: { 'Authorization': `Bearer ${authToken}` }
    });
    console.log('Hotel cargado:', response.data);
    return response.data;
  } catch (error) {
    console.error('Error al obtener los hoteles:', error.response ? error.response.data : error.message);
    throw error;
  }
}

export async function reserva(Data){
  const token = await tokenId();
  const data = {
    "user_id": token,
    "hotel_id": Data.hotel_id,
  }
  console.log("data id: ",data)
  
  return axios.post('http://localhost:8083/reservas/',data, {
    headers: { 'Authorization': `Bearer ${authToken}` }
  })
    .then(response => {
      console.log('Reserva realizada: ', response)
      return response.data
    })
    .catch(error => {
      console.error('Reserva error: ', error)
      throw error;
    });
}

export async function updateReserva(Data){
    const token = await tokenId();
    const data = {
      "user_id": token,
      "hotel_id": Data.id,
    }
    console.log("data id: ",data)
    return axios.put('http://localhost:8083/reservas',data, {
      headers: { 'Authorization': `Bearer ${authToken}` }
    })
      .then(response => {
        console.log('Reserva realizada: ', response)
        return response.data
      })
      .catch(error => {
        console.error('Reserva error: ', error)
        throw error;
      });
  }

  export async function deleteReserva(Data){
    const token = await tokenId();
    const data = {
      "user_id": token,
      "hotel_id": Data.id,
    }
    console.log("data id: ",data)
    return axios.delete('http://localhost:8083/reservas',data, {
      headers: { 'Authorization': `Bearer ${authToken}` }
    })
      .then(response => {
        console.log('Reserva realizada: ', response)
        return response.data
      })
      .catch(error => {
        console.error('Reserva error: ', error)
        throw error;
      });
  }

  export async function search(query, offset, limit) {
    const url = `http://localhost:8084/search?q=${query}&offset=${offset}&limit=${limit}`;
    console.log("Request URL:", url); // Para verificar la URL generada
  
    try {
      const response = await axios.get(url);
      return response.data;
    } catch (error) {
      console.error("error searching:", error);
      throw error;
    }
  }
  
export async function tokenId(){
    const token = localStorage.getItem('token');
    console.log("tokens: ",token);
    const val1 = await axios.get('http://localhost:8080/users/token', {
    headers: {
      'Authorization': token
    }
  });
  const val2 = val1.data.idU
  return val2
}

export async function tokenRole(){
  const token = localStorage.getItem('token');
  console.log("tokens: ",token);
  const val1 = await axios.get('http://localhost:8080/users/token', {
  headers: {
    'Authorization': token
  }
});
console.log("val1: ",val1)
const val2 = val1.data.Adminu
console.log("val2: ",val2)
return val2
}

/*
export async function buscarSuscription(Data) {
  try {
    const token = await tokenId();
    console.log("id user: ", token);
    console.log("id curso: ", Data);

    const url = `http://localhost:8080/suscriptions?user_id=${token}&course_id=${Data}`;

    const response = await axios.get(url, {
      headers: { 'Authorization': `Bearer ${authToken}` }
    });
    console.log('Suscripcion obtenida: ', response.data);
    return response.data;
  } catch (error) {
    console.error('Error en la busqueda de suscripcion: ', error);
    return null
  }
}*/


/*
export function MoreInfo(Data){
  return axios.get(`http://localhost:8080/courses/courseInfo/${Data}`)
  .then(response => {
    console.log('Curso Seleccionado: ', response)
    return response.data
  })
  .catch(error => {
    console.error('curso error: ', error)
    throw error;
  });
}*/
