import classes from "./Login.module.css";
import { useRef, useState } from "react";
import Link from "next/link";
import axios from "axios";
import { useRouter } from "next/router";

const Login = () => {
  const router = useRouter();
  const emailInputRef = useRef();
  const passwordInputRef = useRef();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);

  const handleSubmit = async (event) => {
    event.preventDefault();

    const email = emailInputRef.current.value;
    const password = passwordInputRef.current.value;

    const data = {
      email,
      password,
    };

    setIsSubmitting(true);

    try {
      const response = await axios.post(
        `${process.env.API_ENDPOINT}login`,
        data
      );
      console.log(response.data);

      localStorage.setItem("token", response.data.token);
      router.push("/home");
    } catch (error) {
      setErrorMessage(error.response.data.error || "An error occurred");
    }

    setIsSubmitting(false);
  };

  return (
    <form className={classes.form} onSubmit={handleSubmit}>
      <h3>Sign in to continue</h3>

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
      {errorMessage && <div className={classes.error}>{errorMessage}</div>}
      <div className={classes.footer}>
        <button>Login</button>
        <div>
          <p>
            Already have an account <Link href="/register">Register</Link>
          </p>
        </div>
      </div>
    </form>
  );
};

export default Login;
