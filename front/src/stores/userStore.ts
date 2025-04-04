import { ref } from 'vue';
import { defineStore } from 'pinia';
import type { UserData } from '@/types/UserData'
import fetchUser from '@/lib/fetchUser';

export async function tryLoadUserData() {
  const loaded = loadCache();
  if(!loaded) {
    const fetched = await fetchUser();
  }
}

function loadCache(): boolean {
  const userDataJson = localStorage.getItem("UserData");
  if(!userDataJson) {
    return false;
  }
  const data = JSON.parse(userDataJson) as UserData;
  if(data.username === "" || data.avatarUrl === "") {
    return false;
  }
  const { login } = useUserStore();
  login(data.username, data.avatarUrl, data.repositories);
  return true
}
function saveCache(data : UserData) {
  const json = JSON.stringify(data);
  localStorage.setItem("UserData", json);
}

export const useUserStore = defineStore('userStore', () => {
  const loggedIn = ref(false);
  const username = ref('');
  const avatarUrl = ref('');
  const repositories = ref<string[]>([]);

  function login(usernameParam: string, avatarUrlParam: string, repositoriesParam: string[]) {
    loggedIn.value = true;
    username.value = usernameParam;
    avatarUrl.value = avatarUrlParam;
    repositories.value = repositoriesParam;
    const userData : UserData = {
      username: usernameParam,
      avatarUrl: avatarUrlParam,
      repositories: repositoriesParam
    }
    saveCache(userData);
  }
  return { loggedIn, username, avatarUrl, repositories, login };
});
