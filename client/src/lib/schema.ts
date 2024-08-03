import { z } from "zod";

export const UserFormSchema = z.object({
  email: z.string().email(),
  password: z.string().min(8),
});

export type UserFormValues = z.infer<typeof UserFormSchema>;
