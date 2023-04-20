import { useState } from "react";
import { FaSave } from "react-icons/fa";
import classes from "./SaveButton.module.css";

const SaveButton = (props) => {
  const [saved, setSaved] = useState(props.saved || false);

  const handleSave = async () => {
    await props.onClick();
    setSaved(!saved);
  };

  return (
    <div className={classes.container}>
      <FaSave
        className={classes.saveIcon}
        color={saved ? "purple" : "gray"}
        onClick={handleSave}
      />
    </div>
  );
};

export default SaveButton;
