// Code generated by stringer -i a.out.go -o anames.go -p s390x; DO NOT EDIT.

package s390x

import "cmd/internal/obj"

var Anames = []string{
	obj.A_ARCHSPECIFIC: "ADD",
	"ADDC",
	"ADDE",
	"ADDW",
	"DIVW",
	"DIVWU",
	"DIVD",
	"DIVDU",
	"MODW",
	"MODWU",
	"MODD",
	"MODDU",
	"MULLW",
	"MULLD",
	"MULHD",
	"MULHDU",
	"MLGR",
	"SUB",
	"SUBC",
	"SUBV",
	"SUBE",
	"SUBW",
	"NEG",
	"NEGW",
	"MOVWBR",
	"MOVB",
	"MOVBZ",
	"MOVH",
	"MOVHBR",
	"MOVHZ",
	"MOVW",
	"MOVWZ",
	"MOVD",
	"MOVDBR",
	"MOVDEQ",
	"MOVDGE",
	"MOVDGT",
	"MOVDLE",
	"MOVDLT",
	"MOVDNE",
	"LOCR",
	"LOCGR",
	"FLOGR",
	"POPCNT",
	"AND",
	"ANDW",
	"OR",
	"ORW",
	"XOR",
	"XORW",
	"SLW",
	"SLD",
	"SRW",
	"SRAW",
	"SRD",
	"SRAD",
	"RLL",
	"RLLG",
	"RNSBG",
	"RXSBG",
	"ROSBG",
	"RNSBGT",
	"RXSBGT",
	"ROSBGT",
	"RISBG",
	"RISBGN",
	"RISBGZ",
	"RISBGNZ",
	"RISBHG",
	"RISBLG",
	"RISBHGZ",
	"RISBLGZ",
	"FABS",
	"FADD",
	"FADDS",
	"FCMPO",
	"FCMPU",
	"CEBR",
	"FDIV",
	"FDIVS",
	"FMADD",
	"FMADDS",
	"FMOVD",
	"FMOVS",
	"FMSUB",
	"FMSUBS",
	"FMUL",
	"FMULS",
	"FNABS",
	"FNEG",
	"FNEGS",
	"LCDBR",
	"LEDBR",
	"LDEBR",
	"LPDFR",
	"LNDFR",
	"FSUB",
	"FSUBS",
	"FSQRT",
	"FSQRTS",
	"FIEBR",
	"FIDBR",
	"CPSDR",
	"LTEBR",
	"LTDBR",
	"TCEB",
	"TCDB",
	"LDGR",
	"LGDR",
	"CEFBRA",
	"CDFBRA",
	"CEGBRA",
	"CDGBRA",
	"CFEBRA",
	"CFDBRA",
	"CGEBRA",
	"CGDBRA",
	"CELFBR",
	"CDLFBR",
	"CELGBR",
	"CDLGBR",
	"CLFEBR",
	"CLFDBR",
	"CLGEBR",
	"CLGDBR",
	"CMP",
	"CMPU",
	"CMPW",
	"CMPWU",
	"TMHH",
	"TMHL",
	"TMLH",
	"TMLL",
	"IPM",
	"SPM",
	"CS",
	"CSG",
	"SYNC",
	"BC",
	"BCL",
	"BRC",
	"BEQ",
	"BGE",
	"BGT",
	"BLE",
	"BLT",
	"BLEU",
	"BLTU",
	"BNE",
	"BVC",
	"BVS",
	"SYSCALL",
	"BRCT",
	"BRCTG",
	"CRJ",
	"CGRJ",
	"CLRJ",
	"CLGRJ",
	"CIJ",
	"CGIJ",
	"CLIJ",
	"CLGIJ",
	"CMPBEQ",
	"CMPBGE",
	"CMPBGT",
	"CMPBLE",
	"CMPBLT",
	"CMPBNE",
	"CMPUBEQ",
	"CMPUBGE",
	"CMPUBGT",
	"CMPUBLE",
	"CMPUBLT",
	"CMPUBNE",
	"MVC",
	"MVCIN",
	"MVCLE",
	"CLC",
	"XC",
	"OC",
	"NC",
	"EXRL",
	"LARL",
	"LA",
	"LAY",
	"LAA",
	"LAAG",
	"LAAL",
	"LAALG",
	"LAN",
	"LANG",
	"LAX",
	"LAXG",
	"LAO",
	"LAOG",
	"LMY",
	"LMG",
	"STMY",
	"STMG",
	"STCK",
	"STCKC",
	"STCKE",
	"STCKF",
	"CLEAR",
	"KM",
	"KMC",
	"KLMD",
	"KIMD",
	"KDSA",
	"KMA",
	"KMCTR",
	"VA",
	"VAB",
	"VAH",
	"VAF",
	"VAG",
	"VAQ",
	"VACC",
	"VACCB",
	"VACCH",
	"VACCF",
	"VACCG",
	"VACCQ",
	"VAC",
	"VACQ",
	"VACCC",
	"VACCCQ",
	"VN",
	"VNC",
	"VAVG",
	"VAVGB",
	"VAVGH",
	"VAVGF",
	"VAVGG",
	"VAVGL",
	"VAVGLB",
	"VAVGLH",
	"VAVGLF",
	"VAVGLG",
	"VCKSM",
	"VCEQ",
	"VCEQB",
	"VCEQH",
	"VCEQF",
	"VCEQG",
	"VCEQBS",
	"VCEQHS",
	"VCEQFS",
	"VCEQGS",
	"VCH",
	"VCHB",
	"VCHH",
	"VCHF",
	"VCHG",
	"VCHBS",
	"VCHHS",
	"VCHFS",
	"VCHGS",
	"VCHL",
	"VCHLB",
	"VCHLH",
	"VCHLF",
	"VCHLG",
	"VCHLBS",
	"VCHLHS",
	"VCHLFS",
	"VCHLGS",
	"VCLZ",
	"VCLZB",
	"VCLZH",
	"VCLZF",
	"VCLZG",
	"VCTZ",
	"VCTZB",
	"VCTZH",
	"VCTZF",
	"VCTZG",
	"VEC",
	"VECB",
	"VECH",
	"VECF",
	"VECG",
	"VECL",
	"VECLB",
	"VECLH",
	"VECLF",
	"VECLG",
	"VERIM",
	"VERIMB",
	"VERIMH",
	"VERIMF",
	"VERIMG",
	"VERLL",
	"VERLLB",
	"VERLLH",
	"VERLLF",
	"VERLLG",
	"VERLLV",
	"VERLLVB",
	"VERLLVH",
	"VERLLVF",
	"VERLLVG",
	"VESLV",
	"VESLVB",
	"VESLVH",
	"VESLVF",
	"VESLVG",
	"VESL",
	"VESLB",
	"VESLH",
	"VESLF",
	"VESLG",
	"VESRA",
	"VESRAB",
	"VESRAH",
	"VESRAF",
	"VESRAG",
	"VESRAV",
	"VESRAVB",
	"VESRAVH",
	"VESRAVF",
	"VESRAVG",
	"VESRL",
	"VESRLB",
	"VESRLH",
	"VESRLF",
	"VESRLG",
	"VESRLV",
	"VESRLVB",
	"VESRLVH",
	"VESRLVF",
	"VESRLVG",
	"VX",
	"VFAE",
	"VFAEB",
	"VFAEH",
	"VFAEF",
	"VFAEBS",
	"VFAEHS",
	"VFAEFS",
	"VFAEZB",
	"VFAEZH",
	"VFAEZF",
	"VFAEZBS",
	"VFAEZHS",
	"VFAEZFS",
	"VFEE",
	"VFEEB",
	"VFEEH",
	"VFEEF",
	"VFEEBS",
	"VFEEHS",
	"VFEEFS",
	"VFEEZB",
	"VFEEZH",
	"VFEEZF",
	"VFEEZBS",
	"VFEEZHS",
	"VFEEZFS",
	"VFENE",
	"VFENEB",
	"VFENEH",
	"VFENEF",
	"VFENEBS",
	"VFENEHS",
	"VFENEFS",
	"VFENEZB",
	"VFENEZH",
	"VFENEZF",
	"VFENEZBS",
	"VFENEZHS",
	"VFENEZFS",
	"VFA",
	"VFADB",
	"WFADB",
	"WFK",
	"WFKDB",
	"VFCE",
	"VFCEDB",
	"VFCEDBS",
	"WFCEDB",
	"WFCEDBS",
	"VFCH",
	"VFCHDB",
	"VFCHDBS",
	"WFCHDB",
	"WFCHDBS",
	"VFCHE",
	"VFCHEDB",
	"VFCHEDBS",
	"WFCHEDB",
	"WFCHEDBS",
	"WFC",
	"WFCDB",
	"VCDG",
	"VCDGB",
	"WCDGB",
	"VCDLG",
	"VCDLGB",
	"WCDLGB",
	"VCGD",
	"VCGDB",
	"WCGDB",
	"VCLGD",
	"VCLGDB",
	"WCLGDB",
	"VFD",
	"VFDDB",
	"WFDDB",
	"VLDE",
	"VLDEB",
	"WLDEB",
	"VLED",
	"VLEDB",
	"WLEDB",
	"VFM",
	"VFMDB",
	"WFMDB",
	"VFMA",
	"VFMADB",
	"WFMADB",
	"VFMS",
	"VFMSDB",
	"WFMSDB",
	"VFPSO",
	"VFPSODB",
	"WFPSODB",
	"VFLCDB",
	"WFLCDB",
	"VFLNDB",
	"WFLNDB",
	"VFLPDB",
	"WFLPDB",
	"VFMAXDB",
	"WFMAXDB",
	"VFMAXSB",
	"WFMAXSB",
	"VFMINDB",
	"WFMINDB",
	"VFMINSB",
	"WFMINSB",
	"VFSQ",
	"VFSQDB",
	"WFSQDB",
	"VFS",
	"VFSDB",
	"WFSDB",
	"VFTCI",
	"VFTCIDB",
	"WFTCIDB",
	"VGFM",
	"VGFMB",
	"VGFMH",
	"VGFMF",
	"VGFMG",
	"VGFMA",
	"VGFMAB",
	"VGFMAH",
	"VGFMAF",
	"VGFMAG",
	"VGEF",
	"VGEG",
	"VGBM",
	"VZERO",
	"VONE",
	"VGM",
	"VGMB",
	"VGMH",
	"VGMF",
	"VGMG",
	"VISTR",
	"VISTRB",
	"VISTRH",
	"VISTRF",
	"VISTRBS",
	"VISTRHS",
	"VISTRFS",
	"VL",
	"VLR",
	"VLREP",
	"VLREPB",
	"VLREPH",
	"VLREPF",
	"VLREPG",
	"VLC",
	"VLCB",
	"VLCH",
	"VLCF",
	"VLCG",
	"VLEH",
	"VLEF",
	"VLEG",
	"VLEB",
	"VLEIH",
	"VLEIF",
	"VLEIG",
	"VLEIB",
	"VFI",
	"VFIDB",
	"WFIDB",
	"VLGV",
	"VLGVB",
	"VLGVH",
	"VLGVF",
	"VLGVG",
	"VLLEZ",
	"VLLEZB",
	"VLLEZH",
	"VLLEZF",
	"VLLEZG",
	"VLM",
	"VLP",
	"VLPB",
	"VLPH",
	"VLPF",
	"VLPG",
	"VLBB",
	"VLVG",
	"VLVGB",
	"VLVGH",
	"VLVGF",
	"VLVGG",
	"VLVGP",
	"VLL",
	"VMX",
	"VMXB",
	"VMXH",
	"VMXF",
	"VMXG",
	"VMXL",
	"VMXLB",
	"VMXLH",
	"VMXLF",
	"VMXLG",
	"VMRH",
	"VMRHB",
	"VMRHH",
	"VMRHF",
	"VMRHG",
	"VMRL",
	"VMRLB",
	"VMRLH",
	"VMRLF",
	"VMRLG",
	"VMN",
	"VMNB",
	"VMNH",
	"VMNF",
	"VMNG",
	"VMNL",
	"VMNLB",
	"VMNLH",
	"VMNLF",
	"VMNLG",
	"VMAE",
	"VMAEB",
	"VMAEH",
	"VMAEF",
	"VMAH",
	"VMAHB",
	"VMAHH",
	"VMAHF",
	"VMALE",
	"VMALEB",
	"VMALEH",
	"VMALEF",
	"VMALH",
	"VMALHB",
	"VMALHH",
	"VMALHF",
	"VMALO",
	"VMALOB",
	"VMALOH",
	"VMALOF",
	"VMAL",
	"VMALB",
	"VMALHW",
	"VMALF",
	"VMAO",
	"VMAOB",
	"VMAOH",
	"VMAOF",
	"VME",
	"VMEB",
	"VMEH",
	"VMEF",
	"VMH",
	"VMHB",
	"VMHH",
	"VMHF",
	"VMLE",
	"VMLEB",
	"VMLEH",
	"VMLEF",
	"VMLH",
	"VMLHB",
	"VMLHH",
	"VMLHF",
	"VMLO",
	"VMLOB",
	"VMLOH",
	"VMLOF",
	"VML",
	"VMLB",
	"VMLHW",
	"VMLF",
	"VMO",
	"VMOB",
	"VMOH",
	"VMOF",
	"VNO",
	"VNOT",
	"VO",
	"VPK",
	"VPKH",
	"VPKF",
	"VPKG",
	"VPKLS",
	"VPKLSH",
	"VPKLSF",
	"VPKLSG",
	"VPKLSHS",
	"VPKLSFS",
	"VPKLSGS",
	"VPKS",
	"VPKSH",
	"VPKSF",
	"VPKSG",
	"VPKSHS",
	"VPKSFS",
	"VPKSGS",
	"VPERM",
	"VPDI",
	"VPOPCT",
	"VREP",
	"VREPB",
	"VREPH",
	"VREPF",
	"VREPG",
	"VREPI",
	"VREPIB",
	"VREPIH",
	"VREPIF",
	"VREPIG",
	"VSCEF",
	"VSCEG",
	"VSEL",
	"VSL",
	"VSLB",
	"VSLDB",
	"VSRA",
	"VSRAB",
	"VSRL",
	"VSRLB",
	"VSEG",
	"VSEGB",
	"VSEGH",
	"VSEGF",
	"VST",
	"VSTEH",
	"VSTEF",
	"VSTEG",
	"VSTEB",
	"VSTM",
	"VSTL",
	"VSTRC",
	"VSTRCB",
	"VSTRCH",
	"VSTRCF",
	"VSTRCBS",
	"VSTRCHS",
	"VSTRCFS",
	"VSTRCZB",
	"VSTRCZH",
	"VSTRCZF",
	"VSTRCZBS",
	"VSTRCZHS",
	"VSTRCZFS",
	"VS",
	"VSB",
	"VSH",
	"VSF",
	"VSG",
	"VSQ",
	"VSCBI",
	"VSCBIB",
	"VSCBIH",
	"VSCBIF",
	"VSCBIG",
	"VSCBIQ",
	"VSBCBI",
	"VSBCBIQ",
	"VSBI",
	"VSBIQ",
	"VSUMG",
	"VSUMGH",
	"VSUMGF",
	"VSUMQ",
	"VSUMQF",
	"VSUMQG",
	"VSUM",
	"VSUMB",
	"VSUMH",
	"VTM",
	"VUPH",
	"VUPHB",
	"VUPHH",
	"VUPHF",
	"VUPLH",
	"VUPLHB",
	"VUPLHH",
	"VUPLHF",
	"VUPLL",
	"VUPLLB",
	"VUPLLH",
	"VUPLLF",
	"VUPL",
	"VUPLB",
	"VUPLHW",
	"VUPLF",
	"VMSLG",
	"VMSLEG",
	"VMSLOG",
	"VMSLEOG",
	"NOPH",
	"BYTE",
	"WORD",
	"DWORD",
	"BRRK",
	"LAST",
}
