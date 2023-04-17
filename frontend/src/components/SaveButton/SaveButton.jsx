import { useState } from "react";
import { FaSave } from "react-icons/fa";
import classes from "./SaveButton.module.css";

const SaveButton = () => {
  const [saved, setSaved] = useState(false);

  const handleToggle = () => {
    setSaved(!saved);
  };

  return (
    <div className={classes.container}>
      <FaSave
        className={classes.saveIcon}
        color={saved ? "purple" : "gray"}
        onClick={handleToggle}
      />
    </div>
  );
};

export default SaveButton;
