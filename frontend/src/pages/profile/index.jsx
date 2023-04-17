import { useRouter } from "next/router";
import Image from "next/image";
import { FaEdit } from "react-icons/fa";
import { RiAddBoxLine } from "react-icons/ri";
import classes from "./Profile.module.css";
import ProfileImage from "../../../public/Spaghetti.jpg";
import ProfileNavBar from "../../components/ProfileNavBar/ProfileNavBar";

const Profile = () => {
  const router = useRouter();

  const handleEditAccount = () => {
    router.push("/edit-profile");
  };

  const handleNewRecipe = () => {
    router.push("/create-recipe");
  };

  return (
    <div>
      <div className={classes.profile}>
        <div className={classes.profileLeft}>
          <div className={classes.profileImage}>
            <Image src={ProfileImage} width={150} height={150} alt="Profile" />
          </div>
        </div>
        <div className={classes.profileCenter}>
          <h1 className={classes.profileName}>Malak Shaer</h1>
          <p className={classes.profileBio}>Full Stack Developer</p>
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
