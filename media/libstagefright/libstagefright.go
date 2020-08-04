package libstagefright

import (
        "android/soong/android"
        "android/soong/cc"
)

func init() {
    android.RegisterModuleType("libstagefright_defaults", libstagefrightDefaultsFactory)
}

func libstagefrightDefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, libstagefrightDefaults)
    return module
}

func libstagefrightDefaults(ctx android.LoadHookContext) {
    var cppflags []string
    type props struct {
        Cflags []string
    }

    p := &props{}

    if ctx.Config().VendorConfig("IMXPLUGIN").Bool("BOARD_HAVE_VPU")  {
        var vpu_type string = ctx.Config().VendorConfig("IMXPLUGIN").String("BOARD_VPU_TYPE")
        if vpu_type == "malone" {
            cppflags = append(cppflags, "-DMALONE_VPU")
        } else if vpu_type == "hantro" {
            cppflags = append(cppflags, "-DHANTRO_VPU")
        }
    }
    p.Cflags = cppflags
    ctx.AppendProperties(p)
}
