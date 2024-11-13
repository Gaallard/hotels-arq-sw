
import React, { useState, useEffect } from 'react';
import { FaHome, FaWifi, FaCoffee, FaSwimmingPool, FaParking } from 'react-icons/fa';
import { insertHotel, updateHotel, reserva, getAllHotels, search } from '../../utils/Acciones.js';
import { CgGym } from "react-icons/cg";
import { MdEdit } from 'react-icons/md';
import { FaPlus } from 'react-icons/fa';
import './home.css';
import { useNavigate } from 'react-router-dom';
import { tokenRole, tokenId } from '../../utils/Acciones';
import Swal from 'sweetalert2';

const MisHoteles = () => {
  const [showAddDialog, setShowAddDialog] = useState(false);
  const [showEditDialog, setShowEditDialog] = useState(false);
  const [selectedHotel, setSelectedHotel] = useState(null);
  const [hotels, setHotels] = useState([]);
  const [isAdmin, setRole] = useState('');
  const [mensaje, setMensaje] = useState('');
  const navigate = useNavigate();
  const [reservas, setReservas] = useState('');
  const [searchQuery, setSearchQuery] = useState('')

  const openAddDialog = () => {
    setShowAddDialog(true);
    document.body.style.overflow = 'hidden'; 
};

const closeAddDialog = () => {
    setShowAddDialog(false);
    document.body.style.overflow = 'auto';
};

const openEditDialog = (hotel) => {
  setSelectedHotel(hotel);
  setName(hotel.name || ''); 
  setAddress(hotel.address || '');
  setCountry(hotel.country || '');
  setCity(hotel.city || '');
  setState(hotel.state || '');
  setAmenities(hotel.amenities || []);
  setRating(hotel.rating || '');
  setPrice(hotel.price || '');
  setAvailableRooms(hotel.available_rooms || '');
  setShowEditDialog(true);
};

  const closeEditDialog = () => { setSelectedHotel(null); setShowEditDialog(false); };

  const [name, setName] = useState('');
  const [address, setAddress] = useState(''); 
  const [country, setCountry] = useState(''); 
  const [city, setCity] = useState(''); 
  const [state, setState] = useState(''); 
  const [amenities, setAmenities] = useState([]); 
  const [rating, setRating] = useState(''); 
  const [price, setPrice] = useState(''); 
  const [idhotel, setId] = useState(''); 
  const [available_rooms, setAvailableRooms] = useState('');
  const [cantNoches, setCantNoches] = useState('');

  useEffect(() => {
    const fetchRole = async () => {
      try {
        const role = await tokenRole();
        setRole(role);
        console.log("role: ", role);
      } catch (error) {
        console.error('Error fetching role:', error);
      }
    };
    fetchRole();
  }, []);

  useEffect(() => {
    const fetchHotels = async () => {
      try {
        const hotelsData = await getAllHotels();
        console.log("Hoteles cargados:", hotelsData); // Verifica la estructura de los datos
        setHotels(hotelsData.results || hotelsData); // Ajusta según la estructura
      } catch (error) {
        console.error('Error fetching hotels:', error);
      }
    };
    fetchHotels();
  }, []);

  const handleInsertHotel = async (e) => {
    e.preventDefault();

    const hotelData = { 
        name, 
        address, 
        country, 
        city, 
        state, 
        amenities, 
        rating: parseFloat(rating), // Convertir a número decimal
        price: parseFloat(price),    // Convertir a número decimal
        available_rooms: parseInt(available_rooms, 10) // Convertir a número entero
    };

    try {
        console.log("Datos que se enviarán:", hotelData);
        const newHotel = await insertHotel(hotelData);
        setHotels((prevHotels) => [...prevHotels, newHotel]);
        setMensaje('Hotel creado exitosamente.');
        closeAddDialog();
    } catch (error) {
        setMensaje('Error al crear hotel');
        console.log("Error en handleInsertHotel:", error.response ? error.response.data : error.message);
    }

    window.location.reload();
};

const handleUpdateHotelSubmit = async (e) => {
  e.preventDefault();

  if (!selectedHotel) return;

  const hotelData = {
    name,
    address,
    country,
    city,
    state,
    amenities,
    rating: parseFloat(rating),
    price: parseFloat(price),
    available_rooms: parseInt(available_rooms, 10)
  };

  try {
    console.log("id hotel seleccionado: ", selectedHotel._id)
    const updatedHotel = await updateHotel(selectedHotel._id, hotelData);
    setHotels(hotels.map((hotel) => hotel._id === selectedHotel._id ? updatedHotel : hotel));
    setMensaje('Hotel actualizado con éxito');
    setShowEditDialog(false); // Cierra el modal
  } catch (error) {
    console.error("Error en handleUpdateHotelSubmit:", error);
    setMensaje("Error al actualizar el hotel");
  }

  window.location.reload();
};

  const handleReserva = async (hotelId) => {
    if (!cantNoches || cantNoches <= 0) {
      setMensaje('Por favor, selecciona una cantidad válida de noches.');
      return;
    }
    const reservaData = { hotel_id: hotelId, user_id: await tokenId(), noches: cantNoches };  // Define los datos de la reserva
    try {
      const newReserva = await reserva(reservaData);  // Llama a la función reserva pasando los datos
      setReservas((reservas) => [...reservas, newReserva]);  // Actualiza el estado de reservas con la nueva reserva
      setMensaje('Reserva realizada con éxito');  // Muestra el mensaje de éxito
      Swal.fire({
        icon: 'success',
        title: 'Reserva completada',
        text: '¡Su reserva se ha realizado con éxito!',
        confirmButtonText: 'Aceptar'
      });
      return;
    } catch (error) {
      console.error('Error al realizar la reserva:', error);
      setMensaje('Error al realizar la reserva');  // Muestra el mensaje de error
      Swal.fire({
        icon: 'error',
        title: 'Error en la reserva',
        text: 'Por favor, ingrese una cantidad válida de noches.',
        confirmButtonText: 'Aceptar'
      });
      return;
    }
  };

  const handleSearch = async () => {
    try {
      const searchResults = await search(searchQuery, 0, 10); // Puedes ajustar offset y limit según necesites
      setHotels(searchResults.results || searchResults);
      setMensaje("Resultados para", `${searchQuery}`);
    } catch (error) {
      console.error("Error al buscar hoteles:", error);
      setMensaje("Error al buscar hoteles");
    }
  };
  
  return (
    <div className="contenedor-reserva">
      <h1>Reservar</h1>
      
      <div className="Barra-busqueda">
        <input type="text" placeholder="Busque su hotel aquí" value={searchQuery} onChange={(e) => setSearchQuery(e.target.value)} onKeyDown={(e) => e.key === 'Enter' && handleSearch()}/>
        <div className="date-picker">
          <input className="date-field" type="number" value={cantNoches} onChange={(e) => setCantNoches(e.target.value)} placeholder='Cantidad de Noches' min={1}/>
        </div>

        {isAdmin && (
          <button className="Agregar-Hotel" onClick={openAddDialog}>
          <FaPlus />
        </button>
        )}
      </div>
      <ul className="Grilla-amenities">
  {hotels.length > 0 ? (
    hotels.map((data) => (
      <li key={data.id} className="bloque">
        <img src={data.imageUrl} alt={data.name} className="hotel-imagen" />
        <h2>{data.name}</h2>
        <p>{data.city}</p>
        <div className="amenities">
          <h3>Amenidades:</h3>
          
          <ul className="Lista-amenities">
            {(data.amenities || []).map((amenity, index) => (
              <li key={`${data.id}-${index}`} className="amenities">
                {amenity === 'WiFi' && <FaWifi />}
                {amenity === 'Cafe' && <FaCoffee />}
                {amenity === 'Pileta' && <FaSwimmingPool />}
                {amenity === 'Gimnasio' && <CgGym />}
                {amenity === 'Estacionamiento' && <FaParking />}
                {`${amenity}`}
              </li>
            ))}
          </ul>
        <div>
        <h3>rooms disponibles: <span>{data.available_rooms}</span></h3>
        </div>
        </div>
        <div className="boton-reserva">
          <button onClick={() => handleReserva(data._id)}>Reservar</button>
        </div>
        {isAdmin && (
          <button className="boton-editar" onClick={() => openEditDialog(data)}>
            <MdEdit />
          </button>
        )}
      </li>
    ))
  ) : (
    <p>No tienes hoteles disponibles</p>
  )}
</ul>

{showAddDialog && (
    <div className="modal">
      <form onSubmit={handleInsertHotel}>
        <div className="modal-content">
          <h2>Agregar Nuevo Hotel</h2>
          <input
  type="text"
  value={name}
  onChange={(e) => setName(e.target.value)}
  placeholder="Nombre del Hotel"
/>
<input
  type="text"
  value={address}
  onChange={(e) => setAddress(e.target.value)}
  placeholder="Dirección"
/>
<input
  type="text"
  value={country}
  onChange={(e) => setCountry(e.target.value)}
  placeholder="País"
/>
<input
  type="text"
  value={city}
  onChange={(e) => setCity(e.target.value)}
  placeholder="Ciudad"
/>
<input
  type="text"
  value={state}
  onChange={(e) => setState(e.target.value)}
  placeholder="Estado"
/>
<input
  type="text"
  value={amenities.join(',')}
  onChange={(e) => setAmenities(e.target.value.split(','))}
  placeholder="Amenities"
/>
<input
  type="number"
  value={rating}
  onChange={(e) => setRating(e.target.value)}
  placeholder="Calificación"
/>
<input
  type="text"
  value={price}
  onChange={(e) => setPrice(e.target.value)}
  placeholder="Precio"
/>
<input
  type="number"
  value={available_rooms}
  onChange={(e) => setAvailableRooms(e.target.value)}
  placeholder="Habitaciones Disponibles"
/>

          <button type="submit">Agregar</button>
          <button type="button" onClick={closeAddDialog}>Cancelar</button>
        </div>
      </form>
    </div>
)}

      {showEditDialog && selectedHotel && (
        <div className="modal">
          <form onSubmit={handleUpdateHotelSubmit}>
            <div className="modal-content">
              <h2>Editar Hotel</h2>
              <input type="text" value={name} onChange={(e) => setName(e.target.value)} placeholder="Nombre del Hotel" />
              <input type="text" value={address} onChange={(e) => setAddress(e.target.value)} placeholder="Dirección" />
              <input type="text" value={country} onChange={(e) => setCountry(e.target.value)} placeholder="País" />
              <input type="text" value={city} onChange={(e) => setCity(e.target.value)} placeholder="Ciudad" />
              <input type="text" value={state} onChange={(e) => setState(e.target.value)} placeholder="Estado" />
              <input type="text" value={amenities} onChange={(e) => setAmenities(e.target.value.split(','))} placeholder="Amenities" />
              <input type="number" value={rating} onChange={(e) => setRating(e.target.value)} placeholder="Calificación" />
              <input type="text" value={price} onChange={(e) => setPrice(e.target.value)} placeholder="Precio" />
              <input type="number" value={available_rooms} onChange={(e) => setAvailableRooms(e.target.value)} placeholder="Habitaciones Disponibles" />
              <button onClick={closeEditDialog}>Cancelar</button>
              <button onClick={handleUpdateHotelSubmit}>confirmar</button>
            </div>
          </form>
        </div>
      )}

      <button className="mishoteles" onClick={() => navigate('/mishoteles')}>
       Mis Hoteles
      </button>


    </div>
  ); 
};

export default MisHoteles; 

/*

import React, { useState, useEffect } from 'react';
import { FaHome, FaWifi, FaCoffee, FaSwimmingPool, FaParking } from 'react-icons/fa';
import { insertHotel, updateHotel, reserva, getAllHotels, search } from '../../utils/Acciones.js';
import { CgGym } from "react-icons/cg";
import { MdEdit } from 'react-icons/md';
import { FaPlus } from 'react-icons/fa';
import Swal from 'sweetalert2';
import './home.css';
import { useNavigate } from 'react-router-dom';
import { tokenRole, tokenId } from '../../utils/Acciones';

const MisHoteles = () => {
  const [showAddDialog, setShowAddDialog] = useState(false);
  const [showEditDialog, setShowEditDialog] = useState(false);
  const [selectedHotel, setSelectedHotel] = useState(null);
  const [hotels, setHotels] = useState([]);
  const [isAdmin, setRole] = useState('');
  const [mensaje, setMensaje] = useState('');
  const navigate = useNavigate();
  const [cantNoches, setCantNoches] = useState('');
  const [reservados, setReservados] = useState({});

  const openAddDialog = () => {
    setShowAddDialog(true);
    document.body.style.overflow = 'hidden'; 
  };

  const closeAddDialog = () => {
    setShowAddDialog(false);
    document.body.style.overflow = 'auto';
  };

  const openEditDialog = (hotel) => {
    setSelectedHotel(hotel);
    setShowEditDialog(true);
  };

  const closeEditDialog = () => {
    setSelectedHotel(null);
    setShowEditDialog(false);
  };

  const [name, setName] = useState('');
  const [address, setAddress] = useState(''); 
  const [country, setCountry] = useState(''); 
  const [city, setCity] = useState(''); 
  const [state, setState] = useState(''); 
  const [amenities, setAmenities] = useState([]); 
  const [rating, setRating] = useState(''); 
  const [price, setPrice] = useState(''); 
  const [available_rooms, setAvailableRooms] = useState('');

  useEffect(() => {
    const fetchRole = async () => {
      try {
        const role = await tokenRole();
        setRole(role);
        console.log("role: ", role);
      } catch (error) {
        console.error('Error fetching role:', error);
      }
    };
    fetchRole();
  }, []);

  useEffect(() => {
    const fetchHotels = async () => {
      try {
        const hotelsData = await getAllHotels();
        console.log("Hoteles cargados:", hotelsData);
        setHotels(hotelsData.results || hotelsData);
      } catch (error) {
        console.error('Error fetching hotels:', error);
      }
    };
    fetchHotels();
  }, []);

  const handleReserva = (hotelId) => {
    if (reservados[hotelId]) {
      Swal.fire({
        icon: 'info',
        title: 'Ya reservado',
        text: 'Este hotel ya ha sido reservado.',
        confirmButtonText: 'Aceptar'
      });
      return;
    }

    if (!cantNoches || cantNoches <= 0) {
      Swal.fire({
        icon: 'error',
        title: 'Error en la reserva',
        text: 'Por favor, ingrese una cantidad válida de noches.',
        confirmButtonText: 'Aceptar'
      });
      return;
    }

    setReservados((prev) => ({ ...prev, [hotelId]: true }));

    Swal.fire({
      icon: 'success',
      title: 'Reserva completada',
      text: '¡Su reserva se ha realizado con éxito!',
      confirmButtonText: 'Aceptar'
    });
  };

  return (
    <div className="contenedor-reserva">
      <h1>Reservar</h1>
      
      <div className="Barra-busqueda">
        <input type="text" placeholder="Busque su hotel aquí" value={cantNoches} onChange={(e) => setCantNoches(e.target.value)} />
        {isAdmin && (
          <button className="Agregar-Hotel" onClick={openAddDialog}>
            <FaPlus />
          </button>
        )}
      </div>

      <ul className="Grilla-amenities">
        {hotels.length > 0 ? (
          hotels.map((data) => (
            <li key={data.id} className="bloque">
              <img src={data.imageUrl} alt={data.name} className="hotel-imagen" />
              <h2>{data.name}</h2>
              <p>{data.city}</p>
              <div className="boton-reserva">
                <button onClick={() => handleReserva(data.id)}>Reservar</button>
              </div>
            </li>
          ))
        ) : (
          <p>No tienes hoteles disponibles</p>
        )}
      </ul>

      <button className="mishoteles" onClick={() => navigate('/mishoteles')}>
        Mis Hoteles
      </button>
    </div>
  ); 
};

export default MisHoteles;
*/


