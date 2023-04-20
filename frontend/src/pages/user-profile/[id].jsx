import Image from "next/image";
import classes from "./UserProfile.module.css";
import profileDefault from "../../../public/profileImage.png";
import RecipeList from "../../components/RecipesList/RecipeList";
import { useState, useEffect } from "react";
import axios from "axios";
import { useRouter } from "next/router";

const Profile = () => {
  const router = useRouter();
  const { id } = router.query;
  const [user, setUser] = useState(null);

  useEffect(() => {
    const fetchRecipeData = async () => {
      try {
        const token = localStorage.getItem("token");
        const authHeaders = {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        };
        const response = await axios.get(
          `${process.env.API_ENDPOINT}user/${id}`,
          {
            headers: authHeaders,
          }
        );
        const user = response.data.data;
        console.log(user);
        setUser(user);
      } catch (error) {
        console.log(error);
      }
    };
    fetchRecipeData();
  }, []);

  return (
    <div>
      <div className={classes.profile}>
        <div className={classes.profileLeft}>
          <div className={classes.profileImage}>
            <Image
              src={profileDefault || `data:image/*;base64,${user.profileimage}`}
              width={150}
              height={150}
              alt="Profile"
            />
          </div>
        </div>
        <div className={classes.profileCenter}>
          <h1>{user.username}</h1>
          <p>{user.profilebio}</p>
        </div>
      </div>
      <RecipeList recipes={user.Recipes} />
    </div>
  );
};

export default Profile;
