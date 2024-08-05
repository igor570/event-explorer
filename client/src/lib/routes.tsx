import { createBrowserRouter } from "react-router-dom";
import { Home, SignUp } from "../screens";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/portal",
    element: <SignUp />,
  },
]);
