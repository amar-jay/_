import {
	CarouselView,
	ControlBar,
	FocusLayout,
	FocusLayoutContainer
} from '@livekit/components-react'
interface ConferenceProps {
	type?: string;
}
export const Conference = (props: ConferenceProps) => {
	return (
		<div className='lk-focus-layout-wrapper'>
			{props.type}
			<FocusLayoutContainer>	
				<CarouselView orientation='vertical'
				tracks={[]}
				>
				<FocusLayout/>
					{/* <div>1</div>
					<div>2</div>
					<div>3</div> */}
			{/* <AudioVisualizer 
			 /> */}
				</CarouselView>
			</FocusLayoutContainer>

			<ControlBar variation={'verbose'}/>
		</div>
	)

}