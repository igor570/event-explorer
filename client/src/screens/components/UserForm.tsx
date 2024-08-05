import { useForm } from "react-hook-form";
import { UserFormSchema, UserFormValues } from "../../lib/schema.ts";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { UserFormFooter } from "./UserFormFooter.tsx";
import { useCreateUser, useLoginUser } from "../../lib/hooks.ts";
import { useNavigate } from "react-router-dom";

export const UserForm = () => {
  const navigate = useNavigate();
  const [isLogin, setIsLogin] = useState<boolean>(true);
  const userCreateMutation = useCreateUser();
  const loginUserMutation = useLoginUser();
  const { register, handleSubmit } = useForm<UserFormValues>({
    resolver: zodResolver(UserFormSchema),
  });

  const onSubmit = (data: UserFormValues) => {
    const { email, password } = data;

    if (!isLogin) {
      userCreateMutation.mutate({ email, password });
      setIsLogin(true);
    }

    if (isLogin) {
      loginUserMutation.mutate({ email, password });
      navigate("/");
    }
  };

  return (
    <form
      className="flex flex-col gap-y-8 items-center w-full"
      onSubmit={handleSubmit(onSubmit)}
    >
      <input
        {...register("email")}
        type={"email"}
        placeholder={"Enter email..."}
        className="p-2 rounded-sm w-1/3 lg:w-1/5"
      />
      <input
        {...register("password")}
        type={"password"}
        placeholder={"Enter password..."}
        className="p-2 rounded-sm w-1/3 lg:w-1/5"
      />
      <button
        type="submit"
        value={"Sign up"}
        className="bg-green-500 px-6 py-2 min-w-20 text-zinc-950 font-bold max-w-28 rounded-sm hover:bg-green-600 cursor-pointer transition"
      >
        {isLogin ? "Login" : "Sign Up"}
      </button>
      <UserFormFooter isLogin={isLogin} setIsLogin={setIsLogin} />
    </form>
  );
};
