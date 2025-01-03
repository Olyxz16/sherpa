import { useUserStore } from '@/stores/userStore'
import type { UserData } from '@/types/UserData'

export default async function fetchUser(cookie: string) {
  const userStore = useUserStore();
  try {
    const resp = await fetch("/user", {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${cookie}`
      }
    });
    const json = await resp.json() as string;
    const data = JSON.parse(json) as UserData;
    userStore.login(data.username, data.avatarUrl, data.repositories)
  } catch(e) {
    throw new Error("Error fetching user data")
  }
}
