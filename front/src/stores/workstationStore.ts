import { ref } from 'vue';
import { defineStore } from 'pinia';

export const useWorkstationStore = defineStore('workStationStore', () => {
  const currRepoNameRef = ref('');
  const currFileNameRef = ref('');
  const fileContentRef = ref('');

  const repoName = () => currRepoNameRef.value;
  const fileName = () => currFileNameRef.value;
  const fileContent = () => fileContentRef.value;

  function setCurrentRepo(repoName : string) {
    currRepoNameRef.value = repoName;
    currFileNameRef.value = '.env';
    fileContentRef.value = ''; // Fetch file data here ?
  }
  function setCurrentFile(fileName : string) {
    currFileNameRef.value = fileName;
    fileContentRef.value = ''; // Fetch file data here ?
  }
  function updateFileContent(payload: string | number) {
    fileContentRef.value = payload.toString();
  };
  return { repoName, fileName, fileContent, setCurrentRepo, setCurrentFile, updateFileContent };
});
