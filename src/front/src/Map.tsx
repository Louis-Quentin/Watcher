import React, { useState, useEffect } from 'react';
import { GoogleMap, Marker, useLoadScript } from '@react-google-maps/api';
import styles from './css/Map.module.css'; // Adjust the path as necessary
import MapSettings from './MapSettings';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const containerStyle = {
  width: '100%',
  height: '100vh',
};

const defaultCenter = {
  lat: -3.745,
  lng: -38.523,
};

interface RoundButtonProps {
  onClick: () => void;
}

const RoundButton: React.FC<RoundButtonProps> = ({ onClick }) => {
  return (
    <button className={styles.RoundButton} onClick={onClick}>
      Next &rarr; {/* Right arrow character */}
    </button>
  );
}


const MapComponent: React.FC = () => {
  const navigate = useNavigate();
  const [currentPosition, setCurrentPosition] = useState(defaultCenter);
  const [error, setError] = useState<string | null>(null);
  const { isLoaded, loadError } = useLoadScript({
    googleMapsApiKey: process.env.REACT_APP_GOOGLE_MAPS_API_KEY || '',
  });
  const [retailers, setRetailers] = useState<any[]>([]);


  const RedirectToRetailer = (retailer: any) => {
    navigate(`/retailer/${retailer.place_id}`, { state: { retailer } });
  };

  const requestLocation = () => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          setCurrentPosition({
            lat: position.coords.latitude,
            lng: position.coords.longitude,
          });
          setError(null);
          fetchNearbyRetailers(position.coords.latitude, position.coords.longitude);
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

  const fetchNearbyRetailers = async (latitude: number, longitude: number) => {
    const apiUrl = `https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=${latitude},${longitude}&radius=5000&type=store&keyword=watch&key=${process.env.REACT_APP_GOOGLE_MAPS_API_KEY}`;
    try {
      const response = await axios.get(apiUrl);
      const results = response.data.results;

      // Process each retailer to include photo URL if available
      const processedRetailers = results.map((retailer: any) => {
        if (retailer.photos && retailer.photos.length > 0) {
          const photoReference = retailer.photos[0].photo_reference;
          retailer.photoUrl = `https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference=${photoReference}&key=${process.env.REACT_APP_GOOGLE_MAPS_API_KEY}`;
        } else {
          retailer.photoUrl = null;
        }
        return retailer;
      });

      setRetailers(processedRetailers);
    } catch (error) {
      console.error('Error fetching retailers:', error);
    }
  };

  if (loadError) {
    return <div>Error loading maps</div>;
  }

  if (!isLoaded) {
    return <div>Loading Maps...</div>;
  }

  return (
    <div>
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
      <MapSettings />
      <GoogleMap mapContainerStyle={containerStyle} center={currentPosition} zoom={12}>
        <Marker position={currentPosition} />
        {retailers.map((retailer) => (
          <Marker key={retailer.place_id} position={{ lat: retailer.geometry.location.lat, lng: retailer.geometry.location.lng }} />
        ))}
      </GoogleMap>
      <ul className={styles.storesContainer}>
        {retailers.map((retailer) => (
          <div className={styles.storeCard} key={retailer.place_id}>
            {retailer.photoUrl ? (
              <img className={styles.storeImg} src={retailer.photoUrl} alt={retailer.name} />
            ) : (
              <div className={styles.storeImg}>No image available</div>
            )}
            <div className={styles.storeData}>
              <div className={styles.name}>{retailer.name}</div>
                <div className={styles.storeDataChild}>
                  <div className={styles.childRow}>
                    <div className={styles.square}></div>
                    <div className={`${styles.text}`}>{retailer.vicinity}</div>
                  </div>
                  <div className={styles.childRow}>
                    <div className={styles.square}></div>
                    <div className={styles.text}>{retailer.rating} / 5</div>
                  </div>
                  <div className={styles.childRow}>
                    <div className={styles.square}></div>
                    <div className={styles.text}>{retailer.user_ratings_total} reviews</div>
                  </div>
                  <div className={styles.childRow}>
                    <div className={styles.square}></div>
                    <div className={styles.text}>{retailer.opening_hours.open_now ? 'Currently Open' : 'Currently Closed'}</div>
                  </div>
                  <RoundButton onClick={() => RedirectToRetailer(retailer)}/>
                </div>
            </div>
          </div>
        ))}
      </ul>
    </div>
  );
};

export default MapComponent;
