import React, { useState, useEffect } from 'react';
import { GoogleMap, Marker, useLoadScript } from '@react-google-maps/api';
import styles from './css/Map.module.css'; // Adjust the path as necessary
import MapSettings from "./MapSettings"

const containerStyle = {
  width: '100%',
  height: '100vh',
};

const defaultCenter = {
  lat: -3.745,
  lng: -38.523,
};

const MapComponent: React.FC = () => {
  const [currentPosition, setCurrentPosition] = useState(defaultCenter);
  const [error, setError] = useState<string | null>(null);
  const { isLoaded, loadError } = useLoadScript({
    googleMapsApiKey: process.env.REACT_APP_GOOGLE_MAPS_API_KEY || '',
  });

  const requestLocation = () => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          setCurrentPosition({
            lat: position.coords.latitude,
            lng: position.coords.longitude,
          });
          setError(null);
        },
        (error) => {
          switch (error.code) {
            case error.PERMISSION_DENIED:
              setError('User denied the request for Geolocation.');
              break;
            case error.POSITION_UNAVAILABLE:
              setError('Location information is unavailable.');
              break;
            case error.TIMEOUT:
              setError('The request to get user location timed out.');
              break;
            default:
              setError('An unknown error occurred.');
              break;
          }
        },
        { enableHighAccuracy: true, timeout: 5000, maximumAge: 0 }
      );
    } else {
      setError('Geolocation is not supported by this browser.');
    }
  };

  useEffect(() => {
    requestLocation();
  }, []);

  if (loadError) {
    return <div>Error loading maps</div>;
  }

  if (!isLoaded) {
    return <div>Loading Maps...</div>;
  }

  return (
    <div >
      <div className={styles.container}>
      {error && (
        <div className={styles.errorMessage}>
          {error}
          <button onClick={requestLocation} className={styles.retryButton}>
            Retry
          </button>
        </div>
        
      )}
      </div>
      <MapSettings/>
      <GoogleMap mapContainerStyle={containerStyle} center={currentPosition} zoom={6}>
        <Marker position={currentPosition} />
      </GoogleMap>
      
    </div>
  );
};

export default MapComponent;
