import classes from "./Register.module.css";
import { useRef, useState } from "react";
import { useRouter } from "next/router";
import Link from "next/link";
import axios from "axios";

const Register = ({ token }) => {
  const router = useRouter();
  const userNameInputRef = useRef();
  const emailInputRef = useRef();
  const passwordInputRef = useRef();
  const confirmPasswordInputRef = useRef();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);

  const handleSubmit = async (event) => {
    event.preventDefault();

    const username = userNameInputRef.current.value;
    const email = emailInputRef.current.value;
    const password = passwordInputRef.current.value;
    const confirm_password = confirmPasswordInputRef.current.value;

    const data = {
      username,
      email,
      password,
      confirm_password,
    };

    setIsSubmitting(true);

    try {
      const response = await axios.post(
        `${process.env.API_ENDPOINT}register`,
        data
      );
      console.log(response.data.message);

      localStorage.setItem("token", response.data.token);
      router.push("/home");
    } catch (error) {
      console.error("There was an error:", error);
      setErrorMessage(error.response.data.error || "An error occurred");
    }

    setIsSubmitting(false);
  };

  return (
    <form className={classes.form} onSubmit={handleSubmit}>
      <h3>Sign Up</h3>
      <div className={classes.input}>
        <input
          className={classes.form__input}
          type="text"
          id="userName"
          required
          placeholder="User Name"
          ref={userNameInputRef}
        />
      </div>
      <div className={classes.input}>
        <input
          type="email"
          id="email"
          required
          className={classes.form__input}
          placeholder="Email"
          ref={emailInputRef}
        />
      </div>
      <div className={classes.input}>
        <input
          className={classes.form__input}
          type="password"
          id="password"
          required
          placeholder="Password"
          ref={passwordInputRef}
        />
      </div>
      <div className={classes.input}>
        <input
          className={classes.form__input}
          type="password"
          id="confirmPassword"
          required
          placeholder="Confirm Password"
          ref={confirmPasswordInputRef}
        />
      </div>
      {errorMessage && <div className={classes.error}>{errorMessage}</div>}
      <div className={classes.footer}>
        <button>Register</button>
        <div>
          <p>
            Already have an account <Link href="/login">Login</Link>
          </p>
        </div>
      </div>
    </form>
  );
};

export async function getServerSideProps(context) {
  const token = context.req.cookies.token || null;
  return {
    props: {
      token,
    },
  };
}

export default Register;
