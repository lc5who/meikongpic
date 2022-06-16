<template>
	<view class="">
		
	</view>
</template>

<script>
	export default {
		name:"toolbutton",
		props:{
			ico:{
				type:String,
				default:''
			},
			txt:{
				type:String,
				default:''
			}
		},
		data() {
			return {
				
			};
		}
	}
</script>

<style lang="scss">
	.toolbutton{
		display: flex;
		flex-direction: column;
	}
</style>