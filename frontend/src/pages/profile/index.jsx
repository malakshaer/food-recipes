import Image from "next/image";
import { FaEdit } from "react-icons/fa";
import { RiAddBoxLine } from "react-icons/ri";
import { useRouter } from "next/router";
import axios from "axios";
import ProfileNavBar from "../../components/ProfileNavBar/ProfileNavBar";
import classes from "./Profile.module.css";
import { useEffect, useState } from "react";

const Profile = () => {
  const router = useRouter();
  const [user, setUser] = useState();

  const handleEditAccount = () => {
    router.push("/edit-profile");
  };

  const handleNewRecipe = () => {
    router.push("/create-recipe");
  };

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const token = localStorage.getItem("token");
        const authHeaders = {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        };
        const response = await axios.get(`${process.env.API_ENDPOINT}user`, {
          headers: authHeaders,
        });
        const user = response.data.data;
        setUser(user);
        console.log(user);
      } catch (error) {
        console.log(error);
      }
    };
    fetchUserData();
  }, []);

  return (
    <div>
      <div className={classes.profile}>
        <div className={classes.profileLeft}>
          <div className={classes.profileImage}>
            <img
              src={`data:image/*;base64,${user.profileimage}`}
              alt="profile"
            />
          </div>
        </div>
        <div className={classes.profileCenter}>
          {user ? (
            <>
              <h1 className={classes.profileName}>{user.username}</h1>
              <p className={classes.profileBio}>{user.profilebio}</p>
            </>
          ) : (
            <p>Loading...</p>
          )}
        </div>
        <div className={classes.profileRight}>
          <button className={classes.profileButton} onClick={handleEditAccount}>
            <FaEdit />
            <span>Edit Account</span>
          </button>
          <button className={classes.profileButton} onClick={handleNewRecipe}>
            <RiAddBoxLine />
            <span>Create Recipe</span>
          </button>
        </div>
      </div>
      <ProfileNavBar />
    </div>
  );
};

export default Profile;

export async function getServerSideProps(context) {
  try {
    const token = context.req.headers.authorization?.replace("Bearer ", "");
    const authHeaders = {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    };
    const response = await axios.get(`${process.env.API_ENDPOINT}user`, {
      headers: authHeaders,
    });
    const user = response.data;
    return {
      props: { user },
    };
  } catch (error) {
    console.log(error);
    return {
      props: { user: null },
    };
  }
}
