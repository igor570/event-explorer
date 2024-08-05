export const UserFormFooter = ({
  isLogin,
  setIsLogin,
}: {
  isLogin: boolean;
  setIsLogin: (state: boolean) => void;
}) => {
  return (
    <>
      {!isLogin ? (
        <small className="text-white text-center">
          Already have an account?
          <span
            className="text-blue-300 hover:text-blue-400 cursor-pointer"
            onClick={() => setIsLogin(!isLogin)}
          >
            {" "}
            Login instead.
          </span>
        </small>
      ) : (
        <small className="text-white text-center">
          Need an account?
          <span
            className="text-blue-300 hover:text-blue-400 cursor-pointer"
            onClick={() => setIsLogin(!isLogin)}
          >
            {" "}
            Sign up here
          </span>
        </small>
      )}
    </>
  );
};
