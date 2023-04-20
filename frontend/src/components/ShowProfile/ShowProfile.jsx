import { useRouter } from "next/router";
import classes from "./ShowProfile.module.css";

const ShowProfile = (props) => {
  const router = useRouter();

  const showUserProfile = () => {
    router.push(`/user-profile/${props.id}`);
  };

  return (
    <div className={classes.userInfoButton}>
      <button onClick={showUserProfile}>Show Profile</button>
    </div>
  );
};

export default ShowProfile;
