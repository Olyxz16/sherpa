import { useUserStore } from '@/stores/userStore'
import type { UserData } from '@/types/UserData'

export default async function fetchUser() : Promise<boolean> {
  const userStore = useUserStore();
  try {
    const resp = await fetch("/user", {
      method: "GET",
      credentials: "include",
    });
    const data : UserData = await resp.json() as UserData;
    if(data.username == "") {
      return false;
    }
    userStore.login(data.username, data.avatarUrl, data.repositories)
  } catch(e) {
    return false;
  }
  return true;
}
