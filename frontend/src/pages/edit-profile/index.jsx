import { useState } from "react";
import classes from "./EditAccount.module.css";
import axios from "axios";

const EditAccount = () => {
  const [userName, setUserName] = useState();
  const [profileImage, setProfileImage] = useState();
  const [profileBio, setProfileBio] = useState();
  const [password, setPassword] = useState();
  const [confirmPassword, setConfirmPassword] = useState();
  const [errorMessage, setErrorMessage] = useState();
  const [successMessage, setSuccessMessage] = useState();

  const submitHandler = async (event) => {
    event.preventDefault();
    if (password != confirmPassword) {
      setErrorMessage("Passwords do not match");
      return;
    }

    let userData = {
      username: userName,
      profileimage: profileImage,
      profilebio: profileBio,
      password: password,
      confirmpassword: confirmPassword,
    };

    if (profileImage) {
      userData = { ...userData, profileimage: profileImage };
    }

    const token = localStorage.getItem("token");
    try {
      const response = await axios({
        method: "put",
        url: `${process.env.API_ENDPOINT}user`,
        data: userData,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      setSuccessMessage("Updated successfully");
      setErrorMessage(null);
      console.log(response);
    } catch (error) {
      console.error(error);
      setSuccessMessage(null);
      setErrorMessage(error.response.data.error || "An error occurred");
    }
  };

  return (
    <div className={classes.card}>
      <form className={classes.form} onSubmit={submitHandler}>
        <div className={classes.control}>
          <label htmlFor="name">User Name:</label>
          <input
            type="text"
            id="name"
            value={userName}
            onChange={(e) => setUserName(e.target.value)}
          />
        </div>
        <div className={classes.control}>
          <label htmlFor="profileImage">Profile Image:</label>
          <input
            type="file"
            name="profileImage"
            id="profileImage"
            onChange={(e) => {
              const file = e.target.files[0];
              const reader = new FileReader();
              reader.readAsDataURL(file);
              reader.addEventListener("load", () => {
                const base64String = reader.result.split(",")[1];
                setProfileImage(base64String);
              });
            }}
          />
        </div>
        <div className={classes.control}>
          <label htmlFor="bio">Bio:</label>
          <textarea
            type="text"
            id="bio"
            value={profileBio}
            onChange={(e) => setProfileBio(e.target.value)}
          />
        </div>
        <div className={classes.control}>
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <div className={classes.control}>
          <label htmlFor="confirmPassword">Confirm Password:</label>
          <input
            id="confirmPassword"
            type="password"
            rows="5"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
          />
        </div>
        <div className={classes.actions}>
          <button type="submit">Save Changes</button>
        </div>
        {errorMessage && <p className={classes.error}>{errorMessage}</p>}
        {successMessage && <p className={classes.success}>{successMessage}</p>}
      </form>
    </div>
  );
};

export default EditAccount;
