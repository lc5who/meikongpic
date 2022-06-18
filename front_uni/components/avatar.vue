<template>
	<view class="author_avatar_box" :style="{height:size+'upx',width:size+'upx'}">
		<image :src="src" class="author_avatar_image"></image>
	</view>
</template>

<script>
	export default {
		name:"avatar",
		props:{
			src: {
				type: String,
				default: ''
			},
			size:{
				type:String,
				default:'100'
			}
		},
		data() {
			return {
				
			};
		}
	}
</script>

<style lang="scss">
	.author_avatar_box {
		display: inline-block;
		width: 100upx;
		height: 100upx;
	}
	
	.author_avatar_image {
		width: 100%;
		height: 100%;
		border-radius: 50%;
		border: solid 5upx white;
	}
</style>