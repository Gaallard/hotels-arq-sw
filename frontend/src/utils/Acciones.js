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

export async function insertHotel(Data){
  try {
        const response = await axios.post('http://localhost:8081/hotels', Data, {
            headers: { 'Authorization': `Bearer ${authToken}` }
        });
        console.log('Hotel creado: ', response);
        return response.data;
    } catch (error) {
        console.error('Hotel error: ', error);
        throw error;
    }
}

export async function updateHotel(hotelId, Data) {
    try {
        const response = await axios.put(`http://localhost:8081/hotels/${hotelId}`, Data, {
            headers: { 'Authorization': `Bearer ${authToken}` }
        });
        console.log('Hotel actualizado: ', response);
        return response.data;
    } catch (error) {
        console.error('Error al actualizar el hotel: ', error);
        throw error;
    }
}

export async function getAllHotels() {
  return axios.get("http://localhost:8081/hotels")
  .then(function (response) {
    return response.data;
  })
  .catch(function (error) {
    console.error("error en la carga de los hoteles: ", error);
    throw error;
  })
}

export async function reserva(Data){
  const token = await tokenId();
  const data = {
    "user_id": token,
    "hotel_name": Data,
  }
  console.log("data id: ",data)
  return axios.post('http://localhost:8083/reservas',data, {
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
      "hotel_name": Data,
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
      "hotel_name": Data,
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

  export function search(query, offset, limit){
    return axios.get(`http://localhost:8084/search=${query}&offset=${offset}&limit=${limit}`)
    .then(function (response){
      return response.data
    })
    .catch(function (error) {
      console.error("error searching: ", error);
      throw error;
    })
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
