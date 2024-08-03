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
