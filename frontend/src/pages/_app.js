import "../styles/globals.css";
import Layout from "../components/Layout/Layout";
import { useRouter } from "next/router";

function MyApp({ Component, pageProps }) {
  const router = useRouter();
  const excludeLayoutPaths = ["/login", "/register"];

  const shouldExcludeLayout = excludeLayoutPaths.some(
    (path) => path === router.pathname
  );

  return (
    <>
      {!shouldExcludeLayout && <Layout />}
      <Component {...pageProps} />
    </>
  );
}

export default MyApp;
