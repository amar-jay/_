// store state in local storage using zustand
import { createStore } from 'zustand';
import { persist } from 'zustand/middleware';

interface LocalStorageState {
	identity: string;
	room: string;
	name: string;
	metadata?: string;
	token?: string;
	audioEnabled?: boolean;
	videoEnabled?: boolean;
	audioDeviceId: string;
	videoDeviceId: string;
	setIdentity: (identity: string) => void;
	toggleAudio: () => void;
	toggleVideo: () => void;
}

export const useLocalStorage = createStore(
	persist(
		(set) => ({
			identity: '',
			room: '',
			name: '',
			metadata: '',
			token: '',
			audioEnabled: true,
			videoEnabled: true,
			audioDeviceId: '',
			videoDeviceId: '',
			setIdentity: (identity: string) => set({ identity }),
			toggleAudio: () => set((state: LocalStorageState) => ({ audioEnabled: !state.audioEnabled })),
			toggleVideo: () => set((state: LocalStorageState) => ({ videoEnabled: !state.videoEnabled })),
		}),
		{
			name: 'comrade-user-state',
			getStorage: () => localStorage,
		}
	)
);
