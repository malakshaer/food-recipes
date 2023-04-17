import { useState } from "react";
import { FaHeart } from "react-icons/fa";
import classes from "./LikeButton.module.css";

const LikeButton = () => {
  const [liked, setLiked] = useState(false);

  const handleToggle = () => {
    setLiked(!liked);
  };

  return (
    <div className={classes.container}>
      <FaHeart
        className={classes.heartIcon}
        color={liked ? "red" : "gray"}
        onClick={handleToggle}
      />
    </div>
  );
};

export default LikeButton;
