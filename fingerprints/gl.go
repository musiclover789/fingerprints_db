package fingerprint

import (
	"math/rand"
	"strings"
	"time"
)

/***
列一些显卡信息


*/

var gllist = []string{
	//	`luna_UNMASKED_VENDOR_WEBGL=Imagination Technologies
	//luna_UNMASKED_RENDERER_WEBGL=PowerVR SGX Maca
	//luna_vendor=null
	//luna_GL_VERSION=WebGL 1.0 (OpenGL ES 2.0 Chromium)
	//luna_GL_VENDOR=WebKit
	//luna_GL_RENDERER=WebKit WebGL
	//luna_GL_SupportedExtensions=["ANGLE_instanced_arrays","EXT_blend_minmax","EXT_color_buffer_half_float","EXT_float_blend","EXT_shader_texture_lod","EXT_texture_filter_anisotropic","OES_element_index_uint","OES_fbo_render_mipmap","OES_standard_derivatives","OES_texture_float","OES_texture_half_float","OES_vertex_array_object","WEBGL_color_buffer_float","WEBGL_compressed_texture_astc","WEBGL_compressed_texture_etc","WEBGL_compressed_texture_etc1","WEBGL_compressed_texture_pvrtc","WEBGL_debug_renderer_info","WEBGL_debug_shaders","WEBGL_depth_texture","WEBGL_draw_buffers","WEBGL_lose_context","WEBGL_multi_draw"]
	//luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)
	//`, `--luna_UNMASKED_VENDOR_WEBGL=Qualcomm
	//--luna_UNMASKED_RENDERER_WEBGL=Adreno 650
	//--luna_GL_VERSION=WebGL 1.0 (OpenGL ES 3.1 Adreno 650)
	//--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
	//--luna_GL_VENDOR=Qualcomm
	//--luna_GL_RENDERER=Adreno 650
	//--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 3.0 (OpenGL ES GLSL ES 3.0 Adreno 650)
	//--luna_vendor=null
	//--luna_architecture=
	//--luna_description=
	//--luna_device=
	//--luna_driver=
	//--luna_description=
	//`, `--luna_UNMASKED_VENDOR_WEBGL=ARM
	//--luna_UNMASKED_RENDERER_WEBGL=Mali-G76
	//--luna_GL_VERSION=WebGL 1.0 (OpenGL ES 3.2 Mali-G76)
	//--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
	//--luna_GL_VENDOR=ARM
	//--luna_GL_RENDERER=Mali-G76
	//--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 3.0 (OpenGL ES GLSL ES 3.0 Mali-G76)
	//--luna_vendor=null
	//--luna_architecture=mali-g76
	//--luna_description=
	//--luna_device=
	//--luna_driver=
	//--luna_description=
	//`,
	// 0001
	`--luna_UNMASKED_VENDOR_WEBGL=Advanced Micro Devices, Inc.
--luna_UNMASKED_RENDERER_WEBGL=AMD Radeon RX 5700
--luna_GL_VERSION=WebGL 1.0 (OpenGL ES 3.0 AMD Radeon RX 5700)
--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
--luna_GL_VENDOR=WebKit
--luna_GL_RENDERER=WebKit WebGL
--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)
--luna_vendor=amd
--luna_architecture=rdna-1
--luna_description=
--luna_device=
--luna_driver=
--luna_description=`,
	// 0002
	`--luna_UNMASKED_VENDOR_WEBGL=Advanced Micro Devices, Inc.
--luna_UNMASKED_RENDERER_WEBGL=AMD Radeon RX 6800 XT
--luna_GL_VERSION=WebGL 1.0 (OpenGL ES 3.0 AMD Radeon RX 6800 XT)
--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
--luna_GL_VENDOR=WebKit
--luna_GL_RENDERER=WebKit WebGL
--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)
--luna_vendor=amd
--luna_architecture=rdna-1
--luna_description=
--luna_device=
--luna_driver=
--luna_description=`,
	// 0003
	`--luna_UNMASKED_VENDOR_WEBGL=Advanced Micro Devices, Inc.
--luna_UNMASKED_RENDERER_WEBGL=AMD Radeon RX 5600 XT
--luna_GL_VERSION=WebGL 1.0 (OpenGL ES 3.0 AMD Radeon RX 5600 XT)
--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
--luna_GL_VENDOR=WebKit
--luna_GL_RENDERER=WebKit WebGL
--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)
--luna_vendor=amd
--luna_architecture=rdna-1
--luna_description=
--luna_device=
--luna_driver=
--luna_description=`,
	// 0004
	`--luna_UNMASKED_VENDOR_WEBGL=AMD
--luna_UNMASKED_RENDERER_WEBGL=AMD Radeon RX 6800
--luna_GL_VERSION=WebGL 1.0 (OpenGL ES 3.0 AMD Radeon RX 6800)
--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
--luna_GL_VENDOR=AMD
--luna_GL_RENDERER=AMD Radeon RX 6800
--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)
--luna_vendor=amd
--luna_architecture=rdna-1
--luna_description=
--luna_device=
--luna_driver=
--luna_description=`,
	// 0005
	`--luna_UNMASKED_VENDOR_WEBGL=AMD
--luna_UNMASKED_RENDERER_WEBGL=AMD Radeon Pro WX 3200
--luna_GL_VERSION=WebGL 1.0 (OpenGL ES 3.0 AMD Radeon Pro WX 3200)
--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
--luna_GL_VENDOR=AMD
--luna_GL_RENDERER=AMD Radeon Pro WX 3200
--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)
--luna_vendor=amd
--luna_architecture=rdna-1
--luna_description=
--luna_device=
--luna_driver=
--luna_description=`,
	// 0005
	`--luna_UNMASKED_VENDOR_WEBGL=Advanced Micro Devices, Inc.
--luna_UNMASKED_RENDERER_WEBGL=AMD Radeon RX 6800 XT
--luna_GL_VERSION=WebGL 2.0 (OpenGL ES 3.0 AMD Radeon RX 6800 XT)
--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
--luna_GL_VENDOR=Advanced Micro Devices, Inc.
--luna_GL_RENDERER=AMD Radeon RX 6800 XT
--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 3.00 (OpenGL ES GLSL ES 3.0 AMD Radeon RX 6800 XT)
--luna_vendor=amd
--luna_architecture=rdna-1
--luna_description=
--luna_device=
--luna_driver=
--luna_description=`,
	// 0007
	`--luna_UNMASKED_VENDOR_WEBGL=Advanced Micro Devices, Inc.
--luna_UNMASKED_RENDERER_WEBGL=AMD Radeon RX 5700 XT
--luna_GL_VERSION=WebGL 2.0 (OpenGL ES 3.0 AMD Radeon RX 5700 XT)
--luna_GL_SupportedExtensions=["ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"]
--luna_GL_VENDOR=Advanced Micro Devices, Inc.
--luna_GL_RENDERER=AMD Radeon RX 5700 XT
--luna_GL_SHADING_LANGUAGE_VERSION=WebGL GLSL ES 3.00 (OpenGL ES GLSL ES 3.0 AMD Radeon RX 5700 XT)
--luna_vendor=amd
--luna_architecture=rdna-1
--luna_description=
--luna_device=
--luna_driver=
--luna_description=`,
}

func GetGL() []string {
	if len(gllist) == 0 {
		return nil // 空数组，返回空
	}

	rand.Seed(time.Now().UnixNano())                // 使用当前时间的纳秒部分作为随机种子
	index := rand.Intn(len(gllist))                 // 生成随机索引
	randomValue := strings.TrimSpace(gllist[index]) // 获取随机值并去除两端空格

	// 按照换行符分割字符串，并去除每个元素的两端空格
	splitArray := strings.Split(randomValue, "\n")
	for i, value := range splitArray {
		if len(strings.TrimSpace(value)) == 0 {
			continue
		}
		splitArray[i] = strings.TrimSpace(value)
	}

	return splitArray
}

func GetGLByNum(index int) []string {
	if len(gllist) == 0 {
		return nil // 空数组，返回空
	}
	randomValue := strings.TrimSpace(gllist[index]) // 获取随机值并去除两端空格
	// 按照换行符分割字符串，并去除每个元素的两端空格
	splitArray := strings.Split(randomValue, "\n")
	for i, value := range splitArray {
		if len(strings.TrimSpace(value)) == 0 {
			continue
		}
		splitArray[i] = strings.TrimSpace(value)
	}
	return splitArray
}
