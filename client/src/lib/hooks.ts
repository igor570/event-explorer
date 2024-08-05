import { useMutation, UseMutationResult } from "@tanstack/react-query";
import { User } from "./types";

export const createUser = async (data: User) => {
  const { email, password } = data;
  await fetch("http://localhost:3100/signup", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email: email,
      password: password,
    }),
  });
};

export const useCreateUser = () => {
  return useMutation({
    mutationFn: (data: User) => {
      const { email, password } = data;
      return createUser({ email, password });
    },
  });
};

export const loginUser = async (data: User) => {
  const { email, password } = data;
  const res = await fetch("http://localhost:3100/login", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email: email,
      password: password,
    }),
  });

  const repsonseData = await res.json();

  localStorage.setItem("token", repsonseData.Authorization);

  const test = localStorage.getItem("token");
  console.log("LOCAL STORAGE JWT: ", test);
};

export const useLoginUser = () => {
  return useMutation({
    mutationFn: (data: User) => {
      const { email, password } = data;
      return loginUser({ email, password });
    },
  });
};
