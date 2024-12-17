
import React, { useState, useEffect } from 'react';
import { FaHome } from 'react-icons/fa';
import { insertHotel, updateHotel, reserva, getAllHotels } from '../../utils/Acciones.js';
import { MdEdit } from 'react-icons/md';
import { FaPlus } from 'react-icons/fa';
import Swal from 'sweetalert2';
import './home.css';
import { useNavigate } from 'react-router-dom';
import { tokenRole } from '../../utils/Acciones';

const MisHoteles = () => {

  const appName = process.env.INSTANCE_ID 

  
  const [showAddDialog, setShowAddDialog] = useState(false);
  const [showEditDialog, setShowEditDialog] = useState(false);
  const [selectedHotel, setSelectedHotel] = useState(null);
  const [hotels, setHotels] = useState([]);
  const [mensaje, setMensaje] = useState('');
  const [filteredHotels, setFilteredHotels] = useState([]);
  const [isAdmin, setRole] = useState('');
  const [searchQuery, setSearchQuery] = useState('');

  const navigate = useNavigate();


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

  const closeEditDialog = () => {
    setSelectedHotel(null);
    setShowEditDialog(false);
  };

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
        setFilteredHotels(hotelsData.results || hotelsData);
      } catch (error) {
        console.error('Error fetching hotels:', error);
      }
    };
    fetchHotels();
  }, []);

  const handleSearch = () => {
    if (searchQuery.trim() === '') {
      setFilteredHotels(hotels);
      return;
    }
    const filtered = hotels.filter((hotel) =>
      hotel.name.toLowerCase().includes(searchQuery.toLowerCase())
    );
    setFilteredHotels(filtered);
  };


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

  return (
    <div className="contenedor-reserva">
      <h1>Reservar</h1>

      <div className="Barra-busqueda">
        <input
          type="text"
          placeholder="Buscar hotel"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
        />
        <button onClick={handleSearch}>Buscar</button>
        {isAdmin && (
          <button className="Agregar-Hotel" onClick={openAddDialog}>
            <FaPlus /> Agregar Hotel
          </button>
        )}
      </div>

      <ul className="Grilla-amenities">
        {filteredHotels.length > 0 ? (
          filteredHotels.map((data) => (
            <li key={data.id} className="bloque">
              <h2>{data.name}</h2>
              <p>{data.description || "Una breve descripción del hotel."}</p>
              <button className="boton-detalles" onClick={() => navigate(`/moreinfo/${data.id}`)}>Ver Detalles</button>
            </li>
          ))
        ) : (
          <p>No se encontraron hoteles.</p>
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
      <button className="mishoteles" onClick={() => navigate('/mishoteles')}>
        Mis Hoteles
      </button>

      <button className="contenedores" onClick={() => navigate('/contenedores')}>
        Contenedores
      </button>
    </div>
  );


};

export default MisHoteles;

