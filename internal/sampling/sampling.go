//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package sampling ;import (_f "github.com/unidoc/unipdf/v3/internal/bitwise";_c "github.com/unidoc/unipdf/v3/internal/imageutil";_d "io";);func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _be []uint32 ;_gb :=bitsPerSample ;var _ag uint32 ;var _gg byte ;_ee :=0;_ae :=0;_bc :=0;for _bc < len (data ){if _ee > 0{_bb :=_ee ;if _gb < _bb {_bb =_gb ;};_ag =(_ag <<uint (_bb ))|uint32 (_gg >>uint (8-_bb ));_ee -=_bb ;if _ee > 0{_gg =_gg <<uint (_bb );}else {_gg =0;};_gb -=_bb ;if _gb ==0{_be =append (_be ,_ag );_gb =bitsPerSample ;_ag =0;_ae ++;};}else {_gd :=data [_bc ];_bc ++;_de :=8;if _gb < _de {_de =_gb ;};_ee =8-_de ;_ag =(_ag <<uint (_de ))|uint32 (_gd >>uint (_ee ));if _de < 8{_gg =_gd <<uint (_de );};_gb -=_de ;if _gb ==0{_be =append (_be ,_ag );_gb =bitsPerSample ;_ag =0;_ae ++;};};};for _ee >=bitsPerSample {_dc :=_ee ;if _gb < _dc {_dc =_gb ;};_ag =(_ag <<uint (_dc ))|uint32 (_gg >>uint (8-_dc ));_ee -=_dc ;if _ee > 0{_gg =_gg <<uint (_dc );}else {_gg =0;};_gb -=_dc ;if _gb ==0{_be =append (_be ,_ag );_gb =bitsPerSample ;_ag =0;_ae ++;};};return _be ;};func (_cg *Reader )ReadSamples (samples []uint32 )(_bg error ){for _da :=0;_da < len (samples );_da ++{samples [_da ],_bg =_cg .ReadSample ();if _bg !=nil {return _bg ;};};return nil ;};func (_fc *Reader )ReadSample ()(uint32 ,error ){if _fc ._ac ==_fc ._fd .Height {return 0,_d .EOF ;};_ab ,_g :=_fc ._fdc .ReadBits (byte (_fc ._fd .BitsPerComponent ));if _g !=nil {return 0,_g ;};_fc ._e --;if _fc ._e ==0{_fc ._e =_fc ._fd .ColorComponents ;_fc ._a ++;};if _fc ._a ==_fc ._fd .Width {if _fc ._cc {_fc ._fdc .ConsumeRemainingBits ();};_fc ._a =0;_fc ._ac ++;};return uint32 (_ab ),nil ;};func (_ccd *Writer )WriteSample (sample uint32 )error {if _ ,_dg :=_ccd ._ccg .WriteBits (uint64 (sample ),_ccd ._df .BitsPerComponent );_dg !=nil {return _dg ;};_ccd ._acf --;if _ccd ._acf ==0{_ccd ._acf =_ccd ._df .ColorComponents ;_ccd ._cd ++;};if _ccd ._cd ==_ccd ._df .Width {if _ccd ._dbf {_ccd ._ccg .FinishByte ();};_ccd ._cd =0;};return nil ;};type SampleWriter interface{WriteSample (_bef uint32 )error ;WriteSamples (_eef []uint32 )error ;};func NewReader (img _c .ImageBase )*Reader {return &Reader {_fdc :_f .NewReader (img .Data ),_fd :img ,_e :img .ColorComponents ,_cc :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func NewWriter (img _c .ImageBase )*Writer {return &Writer {_ccg :_f .NewWriterMSB (img .Data ),_df :img ,_acf :img .ColorComponents ,_dbf :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func (_eb *Writer )WriteSamples (samples []uint32 )error {for _ad :=0;_ad < len (samples );_ad ++{if _gf :=_eb .WriteSample (samples [_ad ]);_gf !=nil {return _gf ;};};return nil ;};type Writer struct{_df _c .ImageBase ;_ccg *_f .Writer ;_cd ,_acf int ;_dbf bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _bf []uint32 ;_fb :=bitsPerOutputSample ;var _gbd uint32 ;var _ge uint32 ;_gc :=0;_db :=0;_ef :=0;for _ef < len (data ){if _gc > 0{_gbe :=_gc ;if _fb < _gbe {_gbe =_fb ;};_gbd =(_gbd <<uint (_gbe ))|uint32 (_ge >>uint (bitsPerInputSample -_gbe ));_gc -=_gbe ;if _gc > 0{_ge =_ge <<uint (_gbe );}else {_ge =0;};_fb -=_gbe ;if _fb ==0{_bf =append (_bf ,_gbd );_fb =bitsPerOutputSample ;_gbd =0;_db ++;};}else {_fg :=data [_ef ];_ef ++;_cf :=bitsPerInputSample ;if _fb < _cf {_cf =_fb ;};_gc =bitsPerInputSample -_cf ;_gbd =(_gbd <<uint (_cf ))|uint32 (_fg >>uint (_gc ));if _cf < bitsPerInputSample {_ge =_fg <<uint (_cf );};_fb -=_cf ;if _fb ==0{_bf =append (_bf ,_gbd );_fb =bitsPerOutputSample ;_gbd =0;_db ++;};};};for _gc >=bitsPerOutputSample {_eff :=_gc ;if _fb < _eff {_eff =_fb ;};_gbd =(_gbd <<uint (_eff ))|uint32 (_ge >>uint (bitsPerInputSample -_eff ));_gc -=_eff ;if _gc > 0{_ge =_ge <<uint (_eff );}else {_ge =0;};_fb -=_eff ;if _fb ==0{_bf =append (_bf ,_gbd );_fb =bitsPerOutputSample ;_gbd =0;_db ++;};};if _fb > 0&&_fb < bitsPerOutputSample {_gbd <<=uint (_fb );_bf =append (_bf ,_gbd );};return _bf ;};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_ce []uint32 )error ;};type Reader struct{_fd _c .ImageBase ;_fdc *_f .Reader ;_a ,_ac ,_e int ;_cc bool ;};