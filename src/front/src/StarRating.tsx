// StarRating.tsx
import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faStar as solidStar, faStarHalfAlt as halfStar } from '@fortawesome/free-solid-svg-icons';
import { faStar as regularStar } from '@fortawesome/free-regular-svg-icons';
import styles from "./css/StarRating.module.css"

interface StarRatingProps {
  rating: number;
}

const StarRating: React.FC<StarRatingProps> = ({ rating }) => {
  // Generate an array of stars based on the rating value
  const stars = Array.from({ length: 5 }, (_, index) => {
    if (rating >= index + 1) {
      return <FontAwesomeIcon key={index} icon={solidStar} />;
    } else if (rating >= index + 0.5) {
      return <FontAwesomeIcon key={index} icon={halfStar} />;
    } else {
      return <FontAwesomeIcon key={index} icon={regularStar} />;
    }
  });

  return <div className={styles.starRating}>{stars}</div>;
}

export default StarRating;
