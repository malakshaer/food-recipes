import { useState } from "react";
import classes from "./EditAccount.module.css";

const EditAccount = () => {
  const [userName, setUserName] = useState();
  const [email, setEmail] = useState();
  const [profileImage, setProfileImage] = useState();
  const [bio, setBio] = useState();
  const [password, setPassword] = useState();
  const [confirmPassword, setConfirmPassword] = useState();
  const [error, setError] = useState();

  const submitHandler = (event) => {
    event.preventDefault();

    if (profileImage) {
      UserData = { ...UserData, profileImage };
    }

    let UserData = {
      userName,
      email,
      profileImage,
      bio,
      password,
      confirmPassword,
    };

    console.log(UserData);
  };

  return (
    <div className={classes.card}>
      <form className={classes.form} onSubmit={submitHandler}>
        <div className={classes.control}>
          <label htmlFor="name">User Name:</label>
          <input
            type="text"
            required
            id="name"
            value={userName}
            onChange={(e) => setUserName(e.target.value)}
          />
        </div>
        <div className={classes.control}>
          <label htmlFor="image">Profile Image:</label>
          <input
            type="file"
            id="image"
            value={profileImage}
            onChange={(e) => {
              const file = e.target.files[0];
              const reader = new FileReader();
              reader.readAsDataURL(file);
              reader.onload = () => {
                setProfileImage(reader.result);
              };
            }}
          />
        </div>
        <div className={classes.control}>
          <label htmlFor="email">Email:</label>
          <input
            type="text"
            required
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div className={classes.control}>
          <label htmlFor="bio">Bio:</label>
          <textarea
            type="text"
            required
            id="bio"
            value={bio}
            onChange={(e) => setBio(e.target.value)}
          />
        </div>
        <div className={classes.control}>
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            required
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
            required
            rows="5"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
          ></input>
        </div>

        <div className={classes.actions}>
          <button>Update</button>
        </div>
      </form>
    </div>
  );
};

export default EditAccount;
