<script>
	export default {
		onLaunch: function() {
			console.warn('当前组件仅支持 uni_modules 目录结构 ，请升级 HBuilderX 到 3.1.0 版本以上！')
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style lang="scss">
	/*每个页面公共css */
	@import '@/uni_modules/uni-scss/index.scss';
	/* #ifndef APP-NVUE */
	@import '@/static/customicons.css';
	// 设置整个项目的背景色
	page {
		background-color: #f5f5f5;
	}
	
	* {
		box-sizing: border-box;
	}

	/* #endif */
	
	/* #ifdef MP-WEIXIN || APP-PLUS */ 
	::-webkit-scrollbar{ display: none; }
	/* #endif */
	
	/* #ifdef H5 */
	// uni-scroll-view .uni-scroll-view::-webkit-scrollbar { display: none }
	::-webkit-scrollbar {
	  display: none; /* Chrome Safari */
	}
	/* #endif */
	.example-info {
		font-size: 14px;
		color: #333;
		padding: 10px;
	}
	
</style>
