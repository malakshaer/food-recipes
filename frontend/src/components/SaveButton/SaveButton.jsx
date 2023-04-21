import { useState } from "react";
import { FaBookmark } from "react-icons/fa";
import classes from "./SaveButton.module.css";

const SaveButton = (props) => {
  const [saved, setSaved] = useState(false);

  const handleSave = async () => {
    await props.onClick();
    setSaved(!saved);
  };

  return (
    <div className={classes.container}>
      <FaBookmark
        className={classes.saveIcon}
        color={saved ? "purple" : "gray"}
        onClick={handleSave}
      />
    </div>
  );
};

export default SaveButton;
