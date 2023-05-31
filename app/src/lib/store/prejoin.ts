import { PreJoinProps } from "@livekit/components-react";
import { create, createStore, useStore } from "zustand";

interface UserState {
	prejoinChoices: PreJoinProps,
	userName: string,
	videoDeviceId: string,
	audioDeviceId: string,
	videoEnabled: boolean,
	audioEnabled: boolean
	lang: 'en' | 'tr' | 'fr' | 'es',
}
interface UserAction {
	setUserName: (userName: UserState['userName']) => void,
	toggleVideo: (videoEnabled: UserState['videoEnabled']) => void,
	toggleAudio: (audioEnabled: UserState['audioEnabled']) => void,
	setAudioDeviceId: (audioDeviceId: UserState['audioDeviceId']) => void,
	setVideoDeviceId: (videoDeviceId: UserState['videoDeviceId']) => void,
	setNewUser: (userChoices: UserState) => void
	setLang: (lang: UserState['lang']) => void
}
export const userStore = create<{
	user: UserState,

} & UserAction>((set) => ({
	user: {
	userName: '',
	videoEnabled: false,
	audioEnabled: false,
	videoDeviceId: '',
	audioDeviceId: '',
	prejoinChoices: {},
	lang: 'en',
	},
	setUserName: (userName) => set((state) => ({ user: { ...state.user, userName } })),
	toggleVideo: (videoEnabled) => set((state) => ({ user: { ...state.user, videoEnabled } })),
	toggleAudio: (audioEnabled) => set((state) => ({ user: { ...state.user, audioEnabled } })),
	setAudioDeviceId: (audioDeviceId) => set((state) => ({ user: { ...state.user, audioDeviceId } })),
	setVideoDeviceId: (videoDeviceId) => set((state) => ({ user: { ...state.user, videoDeviceId } })),
	setLang: (lang) => set((state) => ({ user: { ...state.user, lang } })),
	setNewUser: (userChoices) => set((state) => ({ user: { ...state.user, ...userChoices } })),

}))