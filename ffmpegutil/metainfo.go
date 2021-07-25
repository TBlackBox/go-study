//Package ffmpegutil 主要封装图片 视频  音频的基本信息
package ffmpegutil

////////////////////////////////////////meta相关的信息////////////////////////////////////////////////////

//MetaInfo 多媒体数据的基本信息 包含图片，视频，音频 的基本信息
type MetaInfo struct {
	Size   int    //多媒体的大小，指的是存储体积，单位为B （某些系统返回的大小可能为0）
	Format string //多媒体的格式
}

//ImageMetaInfo 图片基本信息
type ImageMetaInfo struct {
	MetaInfo *MetaInfo //基本信息
	Width    int       //图片宽度，单位：px
	Height   int       //图片高度，单位：px
}

//MusicMetaInfo 音频数据的基本信息
type MusicMetaInfo struct {
	MetaInfo   *MetaInfo //音频基本信息
	Duration   int       //音频时长 单位：毫秒
	BitRate    int       //比特率，单位：kb/s  指音频每秒传送（包含）的比特数
	SampleRate int       //采样频率，单位:HZ 只一秒内对声音信号的采样次数
}

//VideoMetaInfo 视频数据基本信息
type VideoMetaInfo struct {
	MetaInfo      *MetaInfo      //基本信息
	Width         int            //视频(帧)宽度, 单位:px
	Height        int            //视频(帧)高度，单位:px
	Duration      int            //视频时长，单位:毫秒
	BitRate       int            //比特率，单位:kb/s,指视频每秒传送（包含）的比特数
	Encoder       string         //编码器
	FrameRate     float          //帧率，单位 FPS(Frame Per Second) 指视频每秒包含的帧数
	MusicMetaInfo *MusicMetaInfo //音频数据的基本信息
}
