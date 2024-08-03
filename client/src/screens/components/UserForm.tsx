import { useForm } from "react-hook-form";
import { UserFormSchema, UserFormValues } from "../../lib/schema.ts";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { UserFormFooter } from "./UserFormFooter.tsx";

export const UserForm = () => {
  const { register, handleSubmit } = useForm<UserFormValues>({
    resolver: zodResolver(UserFormSchema),
  });
  const [isLogin, setIsLogin] = useState<boolean>(false);

  const onSubmit = (data: UserFormValues) => {
    console.log(data);
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
