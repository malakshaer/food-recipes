import { useState } from "react";
import { useRouter } from "next/router";
import classes from "./Logout.module.css";
import axios from "axios";

const Logout = () => {
  const [showPopup, setShowPopup] = useState(false);
  const router = useRouter();

  const handleLogout = async () => {
    try {
      const response = await axios.post(`${process.env.API_ENDPOINT}logout`);
      console.log(response);

      localStorage.removeItem("token");
      router.push("/login");
    } catch (error) {
      console.error("Logout error:", error);
    }
  };

  return (
    <>
      <button className={classes.button} onClick={() => setShowPopup(true)}>
        Logout
      </button>
      {showPopup && (
        <>
          <div className={classes.overlay}></div>
          <div className={classes.popup}>
            <p>Are you sure you want to log out?</p>
            <button className={classes.yesButton} onClick={handleLogout}>
              Yes
            </button>
            <button
              className={classes.noButton}
              onClick={() => setShowPopup(false)}
            >
              No
            </button>
          </div>
        </>
      )}
    </>
  );
};

export default Logout;
