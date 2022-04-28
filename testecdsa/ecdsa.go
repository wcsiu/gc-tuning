package testecdsa

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
)

var PK *ecdsa.PrivateKey

type Sign struct {
}

func (s Sign) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := `MEUCIDFqeoSJN7mtat3MjmFGe6NrSNv0GWK4/j4+0RrUU8wpAiEA0piwztKqZLlCDP8QCtgP5oCCSil+HtnRYRlel106tPo=
	MEYCIQDWgcE947qbZWw2QetuPCGZUxKovrTLvmVP/ctPYAzRBwIhAPi04Nc9NuOME6loTgatZZxpOhnlCNUqGDxGcKNWcWys
	MEQCIEXxZ9F0uS1/xcgWhoElD2AZjiSa+mwtjFoFAVhM/N31AiBQiK095yVivYi9cFsk4pwZYGNLWRcjlA+aI0CK5Ul8fA==
	MEQCIEoTgjfQZWqM/cv1Um5PHcB0bSDQn6czh0jT++g3ZpDNAiA4Wv9+JFfmYiBdKS9gTB6OqppMEUuxsa3guatAS93ALA==
	MEUCIDvzBxvlmdw6EHZryCpnaXGcINSO/7nsrN1hdGDOQmAKAiEAm+6xy5TpAunf0+9/8xYXkoKB6CuEczTW/qCkKEFGKAA=
	MEYCIQDoi5FvwH3DxoZYcNpeFzn5P/V8etlNXU3Ti1LODEUvkgIhAI9tGVvVPz+RXuSAxtKo/cFkGMMu76OSn1FtrnJHJL8D
	MEQCIG5zDd3oVkm0L4vWP9lLoEtlQGq9kRwR0lqUKeKVvD4oAiBcvPUR3QFKvWIk3CjuoZ8f4YMTK5c798tLXo2baqszjg==
	MEYCIQD+FtTkL8+jBC38B2rDGKxI1qK/bjZc8YcRnnQbl8hLvwIhAKT7hS6w+dYD0Iqh5snRPrytEGtBCPB3Qif44i46y7I5
	MEQCIGDAzmqkyFiu73+S41rLjjM1gDwykqhI1/HqVdVV1+7xAiAJ0ZLI73qq6BFqkkTb+305WWF8iR/i/1269bJAUwYnNg==
	MEYCIQCgbqmvnjb2qMrQ0uLQkTknCPBFxvllY/di2/oQU7gCjAIhAJUx05zOJSbTvGgd//0qehmELJsn+zukTRPr2uPduYsp
	MEYCIQClp3rNlAjQH7mRs3ktMCft4CSB37gQmB09+IxFEKCONgIhAOEI4YqR35o1ubiOs4/k9YQdCPvSb0TGAegvXBze6i/U
	MEUCIQDIRXwKBpIIFxPN6Zg4NXvcgMTvdf/tlxpIpGIsbdVxhgIgDgSljNqHqze2QvUmBO4SVBMakbO21UtySOvUx0hy4B8=
	MEUCICNG0xjVMEVAAUBMjtj/7oi3JceImmMFiDhcV7s+DV+vAiEA/9ygIbWZi9pfNAkDHM06Sz/y+gaj78jV/OStP0I/uSU=
	MEQCIC0ymQ/8pyjotZJFD7bCV1W6KlUKnAPfmXbQJNIQObvQAiBLEvcRD3U1i5rBqhvmMFP4p7bdsJQTQhpXOju6jFn9lQ==
	MEYCIQCcnQ23xk3GzxjjF2Ji/MVwcRCi6XM9ZB0JKr8YlPLTeAIhAIZOH75PmUhn9pmlnAlxNEkarvFsD9PwftnkIvVPQPJu
	MEQCIHt8aIIApZSrSh8V9cQQSpNrL8r0hqB7HCeh/hWS1t3BAiAiPTwyqI4e+FeirU5CqQUXOdW/5YxvEY6D5/qfCdaf6Q==
	MEYCIQDFAY762pVoRCeZBlm5vBzFsmlpASGIfz9u6EsXDlsrmgIhAKG8TujG9NOkpSUWOuErtqBtp+oCvfUiLjTu14sIKwGq
	MEUCIQDXv0rstoO8tV5Kdr+o4XKaqqpgcrxgwDOa31KG0NzsZgIgYSx6j5+WwMSg0j5NA/20jy6bM1fS08azSUvmHvKq62c=
	MEYCIQCZrpWVV8XbPeueXFWeIqfPqnqWpQxR4Iys1DtK3qdOHQIhAOSmDdT+N9+YHZVfjD71STYRqKAdnHmR5tAhOjunxBq6
	MEYCIQC/dQ4yn0ADI7oFhA+/c3AqduXJ2Ei50CgSH08ZpvRIggIhAIIECyUx1wq5rPcAqMDPgUYNs0tm84jmdkvEMJ0nBmlT
	MEUCIEn/fMGPgUWup8Qm1rBClBhLINx36+WwqkSougcHezhIAiEAu8LFTCNgojd815XjGNq79LQ4IFsQDXJz5J5beDyJplM=
	MEYCIQCEA1bwm5YtRa6iNgdun3UuIS3me+Tk8VUX+SYiJgwplAIhAJfMQOHmSY/qlZ+2fss7gMiHfW8cNpkjnO51/8HdJYWT
	MEUCIQClZW6RewKMsa6xyeXTnhUXINf89PNBQpD2HWMtOHz3mgIgQH7Daqd1FfkBjuUgqTbyiBr7xMnjEN7RHlfUeH7Q4/c=
	MEQCIA4UnX0YstsR2IYo/2cKTF9sgq+vPP7rZWb0L4zvyprJAiBAM5gNmssIaDPNSxFBQ3xSKt3n1CIxdtjR7xmT7YX8IQ==
	MEQCIEDPLkVb1wqz/tRyKRvt+U9hls84QpPIhO7tcTTGbY3zAiALZDJ1i43yDTpJlkC13Xtua60RRtmBwR5IqOOIhxxVIQ==
	MEUCIQC+F/ufJnax8qnAhksGdzme+eJRABihq3FALPbhzs16lgIgC3xToqeBwkBP4m/vvhoZJTBOfbDRdB57hfkEuRHSO3M=
	MEQCIHw3ZQxCDdRrP6gzwWYhbLL/LoIQwzwAYkol9i5Xx7nxAiAI3sxBsWCjiCeFhED+r72RgkE0kDO0O4k/KTrJPCeU+A==
	MEYCIQDv2KW+sT264wCv4ia4sZLQHCxFoqWh+XVyIkbkIIsm6gIhAM+ectwGDC7Kta1UyoUCIO5UT8XV5My258Cc3a5tccNB
	MEQCIENERZDhlDTYvCKH1Kr5QudFFtzO2UFOob+8CGgb4B/FAiBY9t3I+7mJJzKx13qBoyepwY3M23qvZemFWFSzcfYA6g==
	MEUCIFntvdOjX6gH3p91FpZVyua9OdexT0AJbTXmoN0cB62BAiEAk1EPtgBkJfE+xDSz+wbw/4i7+XQXex9r0Qce06DHlrc=
	MEYCIQDLaftxKXklStTWxxnTcvV+Ai9rOyIuqsyK2002AUAUvQIhAP1JbJzQv0mTRi5J8sQIXpm2z2RmpbSD+G9lE2Z5ZgVG
	MEYCIQDnzCv/yYmOuuDCMDxaAfhVhsbg3TiC37/2nd0rLPtgYgIhANxs3P+b+PndQ+1GK3TGuQjHPHI/QfsaZtUcJgHus9R3
	MEUCIQDUuKdQ5C2VgQv7hPCLS71qS+M+XltyKhMoT/3CqkK+0QIge8qYBUEubTUYxAqFf59L2PjlD4KlYBP6tqm33ddyA88=
	MEUCIQDtkfdvypCyOUxDPF1kJkwlEjxVAs5Tz8cheQ79S0qHhwIgQqSVbHQa6wUjNR6O+y4jXE6+uZEwDZpqL01sqW5u1zs=
	MEUCIGQ/WPXbrbxu3qlAQk9hRiFYaUHV4F12JBNYmG5milgSAiEA8GTw4Z2btFw3BTf7zJY5QjnoC+cADKeo+ciSAA7LaMk=
	MEUCICHlKfOdFpdEpIhtbIBTM5g8/e3Ki5qMg6XJ5k30Tf+qAiEAwLef5sfCq/82H8dTDdBY9C05MBUyCuLYzMvQJ/yrLTQ=
	MEUCIBnayzWILOMe+ENLggNgocA/iwplGBatbRkQGStEM4KVAiEAnMK+zrdLIumQ70p/P+GTS2VYnUEFxsGQfw9ectwOMVY=
	MEYCIQDVIcgcnANNQxoJyFiOq5khVMNEnxEbo0uvV7gkcKiP7wIhAK3XA3f20Go+Zt9Mj+F/2j1qAHaLfxBJFXwFZ23D6GY+
	MEUCIDpYZQ99hgszNyZ+ccOyPbgFDi9glZC7IVywp6bWzLzoAiEA41WO/mY3zH2biOHw5TjLP20EBhJQNnuJZXhsvL00A6A=
	MEYCIQC2O0vmctzkBho+cSLf1zWOSDWRgs+DlV/nxg9xwyyKywIhAMDpjvkDh9DRg05+eaqebQCmdkdhEoE59XLK0rkAzCFW
	MEUCID6PFCLXqPS9ptHNalQvejrNfofs+XJ4K8B7iuqjUWvwAiEAt1goCKgrV+t9bHV53MFBECv6N+DGJ6MBkXTtD6PvVRw=
	MEQCIHkjpwbnpWxRdcUYhPNflcVF/W1MylGl0afJBTdVZE/SAiAFnGa4XcqoEqpK0ssATDRtw739qO1GZoOkb8e29cLuYg==
	MEUCIQCwKIH8EH1lwJqLGUcNu7gNpkPbJ2X6nAoggmaiCchPkAIgUWQcrb/GZuhjHLTBsIcRi6dG5X4tV+95MQuZlD9hxuI=
	MEYCIQC2ghLGbj+iwGD3wSY/jxmvsmj5XOugdHoH7nee4HjzmAIhAPS0LH+MrvrzKjOFH3Ef5rTfaQhwOUoYBKUIi7eL/Qj2
	MEUCIQCDzt1COsmRL72S3MmmTtS0z21Yvn8fwfJ6gsfutxO46QIgS+rIVqg5ITa44GIc21kT2lK/u2WbIi5M8xKmmuj5CQ4=
	MEQCIFzq9Ch5rnjwVdGrhXRN88nzfRrpGPEG2yAGQR1E6zc2AiBqj4NE9mwDSNC8NZ4yLB6o9MTkDMmQWKRcXQSJSWDY3g==
	MEYCIQCOHm6V/c5bJo0BYMpzHdEsJDO6cjlGBQq9Z5+yfw236QIhAKOQH6s9ujEqSDCcyccMT72Wi3wh+ZIU/oIPIA+umN3A
	MEUCIQDxN+0n/rvDEf/b1DUaTTaARz9o8K5ZMGga8krnX1ANIgIgey1nNLpMhgXmXXXCHt69AJ4UgPJIep6l1BTooekfc7A=
	MEUCIBZmea0kihauspmKKu9MElOiznWO0BowynX+h38ePO3PAiEAjUtPw/JYNRi38/Gu9NZUolLRgIC725X0sKKd6O/bVZM=
	MEYCIQCq5dRU8Y67eZscJDkGNSJKMlMqc4l8awDYDyr2O80E6AIhAM/wZW1qmsy13i439yzZDgcrIWTKXUInBEOVbP00GSXQ
	MEUCIQCvv3s9ds+GpDamK8Tk07oCVs4Mn+YReF51gj/H2nprewIgDsChJy4gWkGc8g2XVbWXwymV9e8o9s/dbZ/QG0yFagw=
	MEUCIQDJUtbcVY3968pqUuSpvOxR6F8aMxv3eEqqzaX3VoWDpQIgM5hAETTZ8vLE8hBiaNO7NaMRtzZDcEXMrLMCt+1eSOU=
	MEUCIQC0FSOp9VxoXonG+K6pkHc8j71ex/Zx5qzEgoO8Sj5QvwIgQX+8s65IRC1gzTEU2mRrQIaFDYeq0iYmBWlm7nfQTUs=
	MEUCIQCyrZr2/58Dq6nimTgRQ07f4bS7zK1kXTA8OG81zWyS4QIgMckfegnpBexIvQFReU6QAncx29nlFXbi84Pjw89fhtQ=
	MEUCIG7p8NvhKaysf1PIK7ixXzgAbvtrkLzFTXQYL5zEcOnSAiEAp1Gyi0tLkBB5r8axX8Tyc4erIMwNt/JohUfcSG9vixc=
	MEUCIQCfaIfJYLGCEWi/QpR3F9R64WKIPb7pgCXlv8/m3DtCpwIgZgusiocyJZnx/EygQ05fOCLMfGFsAEhCdNooHuU1E6Q=
	MEYCIQDbPPeDgyhSQU1YQwG/wXLueL5pMQfLIecY9UcLJfyKzwIhAPvfQICI/fzE8rOtlJ4upvoG67yu6FYsI8xrQ3aWjcLY
	MEQCIC59h+5hEWfNq1q/guMnK6ZvbsEuRBcduNKvRrjNl8otAiByjVeLss53xPXpj+Bs6RhKmztCs73Ueeg/9PFljeEHkw==
	MEUCIQD3/R6ULT44ANLBLIm75D9qC74t+YHsm7XdnaYHdJcu9AIgGhjTHkhLImkAdcccq0UR9i0EtzrypZ931AEBxvIJ8wg=
	MEQCIF4Bt55R2N1lgLaaCMnZ3nENnNc9hvVBkF0iiWtpjQJgAiABsVXQeofNOMnpGZCPAtQkT/JnfPJxBZPzgr8JxEho+w==
	MEYCIQCgQ5NrgaW+CW4m/sncjP4ZCKJIAl7kiSAy1RTcvrQW7QIhAJ59BW+97vmyDedtn4YGgcIOovpN8ccPySNWM3OcAGC0
	MEYCIQCj1jx0qUBf/M40MZIyjG4+SSrYGdfCGGpoWHq7qmhGiwIhAKxl7D5wkO5dpqCo/8TQiTOZ4XHievcmIlmUkjEFnk70
	MEQCIHgarxzViKcKe7/V0xzg50YVufS21J8eh27onOopxGfpAiAMPn71HnzUh4YB5/2rE2SkY3CWhVLOSa92eqIXvNWgGQ==
	MEQCIH5qFO5EuzF+BxyV10avex2USgVGMm8Ctc1XVNNK/w0+AiAkWIaglINSdhxpS/TvsRzoVPx8giMY/B7IhD9jaQRzpQ==
	MEYCIQDwa3CH1b9zi84cjZT0OJ4dg2Ao748BnzPKdIsz9CEm6wIhAJMMvZQizo9FrNOjlobwJhWblkoieKj6+NnyIM/GzACp
	MEUCIQC7EFScFAGdVl4hxaO86IA4i2W6SBWRWIv+SP2bTPJsoQIgFa2XdFO0dJ8ucF26tMU1sNqAH22wQQ+XZtKh09Tt/I0=
	MEUCIQCCKuTN5jR4aKAxjmKZb6abtgElT4vubt5bPfPfl5xn6QIgHRoj7vdGOzqwC/ZJPR3K0jdFgsQCVZt6yzWJJe7Xslw=
	MEYCIQDnArTqT3wIkExbUcPRTx96I0G2DHcJKOdtLAelmvEsCgIhAOoB9R9tcX3IbXyVrtxiCKTpy0O3wkgsrUB/5/66n1TL
	MEQCIAuH3FIkkDBo4+3U1vwYogpRtB0Td/53g/cFh8Tv88+NAiAitf+8oaHagnesWiCGIKueXJ9ePoEghLA9kYZ0Ke1oEg==
	MEQCIF/1r13NxFTmJvWl71JIHFA7ncD0evNI5FzFiNNtBd+XAiBrdSEXXgN3irlntZFhzpNu9/9QDhgeEYhT0HorEfjiBg==
	MEUCIQCFoa3Nsv2vHTV23rlJYpAMb4h53Jq5s8HdDUShc8IumAIge6b2rqGeIMqoljE4Icp59afT8Ft/WCMP+g60NUpBEhI=
	MEUCIH9ThApIWDiHA2JldGF590R4RIGFYA0VpA3eZ/M81T0DAiEAg2sN9gqq9odukkL4MFQlFllFXkI/85FVuNvVn+zKr6g=
	MEUCIBV7zEFRAa+jj2zeDAhn9bcIySJiqbF1Qu+Jo6Ft9z9jAiEA0BAJ04wK1kWOhnnrp3To7YsVYEuAwsXste2+c7IEQw4=
	MEYCIQCgccnawxWhsT2ljVqOwieRALJJED2yMOFZYvqXfJ6zgAIhAJwTv5brBcgwRSClwKnJkpWwi1L4jaWJePARFKvc1jdS
	MEUCIQDurmx8RkFDpzan9tXd6RQyy9R1N5/zvuWzajMSODyqQQIgDq1FSdu1TjsCR2aIiHVbJSCFYi9wsZw4zayzF7DiBjI=
	MEYCIQCZ3lcWlF0OLiobLIdQGB9WUrZp6T85lWjdNeWrqzx/dgIhAI3k61/wERqLqSgq6Z9Gt2QkD//pq0aKdxSL33fAYbQD
	MEUCIQC77qXlIaD3VPZMgHPeK6jyDRsJgOVIQJMeW+84EOQVpwIgJvfCJJACML1vRg/468VhkxhSEIHjxIfWtUY82k3GGgM=
	MEUCIQDy95JUiTiP6sdgndQyrNnbHUAigocPIxgxCYbUbnXePQIgKbLi2BVnsWKYbXUKrJPhoxY69yqHGVKZrvoDoHfETfM=
	MEUCIFPoQUuhelX10AdlO2jNZGETAIHF5duzC6HVv0lVOexRAiEAyIlLVzw0QCoUDQkeP0EtExqJsTSuxV6LodJT6Qqs+XE=
	MEQCIHYm3WAM8NLywBgFZKuBdy/JRJzkM2CJOVfZN0K3f0PrAiAffgBOnaBcTTUtuGX6KwwgJWwz1OEQak2oXMXcqmfH3Q==
	MEUCIFVzLhcLgK3s4WayP+j6ibzVjRNxI4k8y1pkeV+HY1vqAiEAqhVGOM2fNt7L2TGaXMCfds7wC1dyRax1MYPGWg4pKzs=
	MEUCIQCOWrpXA8KZ4O/axvwSAwJWhHK+EQHr7qEmECdj7orKWAIgMaHd1pWWqPAqwvKvEFFlgC5ujnKmKgmHYF9/g7WV2NA=
	MEYCIQDjbGoCoOWA+TKFzeQvuVJSJTKjXqLv4lFeCt1v91NL6QIhAO5EMNnQ+dEhi6tr0C2l9azx5CaSto+MXeb6cMrSZ7ea
	MEQCICO/G4DqgVYXu5//+bHRXh0suaqsyKeYcTNYedv2EtviAiAWpitJTK38lPAW0q9zvsbtvh2dRNpOwnhWRPZTJ2LtyQ==
	MEUCIQCdxfZ4iKnh4/UMkr5pt2zSh2NuDfeLJRsu9f4+0GiVNgIgDkuxPj/TtKyY6QSNW41WwGUy6OgbXkNUo1zxpnvWBKg=
	MEQCIHhAeR0PyxRn7JJrJO3BoVZSU/ESSbXmTX8DeYbWOKG2AiBd/GVJE2rL85fANXk5S+/LetCXZrpT8jqRLeieFu4Clg==
	MEUCIQCaLx7i1fL0tpikhxJMEndYRVys1mpQ3YOWEbVL8x1JnwIgNXYCqmRZvVxAO8SWrhH6YPnAAcomxUL/iBZxGOdQO90=
	MEUCIQDnHfvIKfun1dCqSio0PtWLO1fMAsDujIB3RXkhP1b/cwIgFYnQRg4M3T8tELytQNU/Ypra8JnMTxDimAU7rcL5ySs=
	MEYCIQCivZUFCfbo08umMNi9i0xirxFjvHBoRqwhrP3WYLviTAIhAJHB1Qm2nqAfOYSCCEa0WzecNKLxy3eZw9Bu314NOyip
	MEYCIQDvxmbR72eXHY21FXEoILfnnJNm25uux26IVxYkUdTemwIhAL027Arq9s13Q0VswYCeerrOMH47+LjP1ogpp+9l0ZcU
	MEUCIQDajG+G9W+QPJ+R5arcYnS6sGOVeFLTvRXeECx/W+cjBQIgdLfGJku4OxCKZFLOBQfkpzqtiFiaOG37KySH/xq9Fa0=
	MEQCIH+P7bC/2GuFepp1egrKAI6AvBfFxb/xV7654hzKDoSzAiAZMI+3wLKf+F/P+BaRr5U87qi2NiA09whhlrnw05TO5Q==
	MEUCIQC5HXJkqdW2M1eFQhvJFOrZC3cOBdpkgxoXnQwo1CmA9gIgE/9eYChPp3qlKYhv4qfNvEtoMXfP1ICqt1cVXWYkG8Y=
	MEUCIQCn0X2T4sAxF6aBmxO8tLqI9P1lO1JbncGQy01WhoYgaAIgAzjxRqObpgjibYhheG+mjA7UY010wgAISK+xMKuSmiQ=
	MEUCIDE8iwyWTSQEpjk5WFwajZgFvmpYnvCREmYWKJmE8S9GAiEA7XjtSw81r9O+wzNy1PNLn22uzhh669fizkpwVsdCsGY=
	MEUCIQDhXwTCn2ZZR+b9FF8hNhGj9V/b7DaXYI8/6S0x9Jd7KAIgDnel1dKeMNR6iZ8h77jlp/Ef4vBdad4hb2TMlZyVOtw=
	MEQCIBwxpOQQAE7TQ/KW6y7+UuGZaEZmxu7d1Lj7K69DfYQIAiAxyQ1/tm3AUJnntAIpIWluIW0DIzTTkusu81MW9Sm8VQ==
	MEQCIBBxGef+Ji2NRrHAMoG5iig2WnclCV/df3XYO4fysbD0AiB5FCKROT/Pyxw957Oz7M6qLqla2vzeZoJgzZSU9vrasA==
	MEYCIQDH8nNTxSMJEOmE2agmnAET/DB+gZLWETkqejqDrGLlUgIhAMM41PqRPXOm6S6duU9VRsN6LaS0v+jamUVDRMT2wxXx
	MEQCIHHsxqsPqy2JP5oGchDwhrbaLqD4eTObU+1UPCQeqOaVAiA9v5BmAC4FlZVS6+liZQCZEWDNP+tbm6GqrahUb2ID6w==
	MEUCIQDFDRXmNW/0TFffgbbDyaRxYq3IaFhgV1OLyJp8Y1Y5eAIgDlZ6fH9PAWBFBVz2fbESp+i8bW0hPSEjoM0juU1BGi4=
	MEQCIBSW0ykHrFKm0QffrZx80HKFoxi3o9DRNWwiwbACTJmGAiBNh7U5lwE1aVWbWcdriAzs8nuIkfVl58OX3RI3pBEvWw==
	MEQCICJvXX20fKJOBo0TJ76xbH6OaCbgz9vBc4Kihdkkp8xyAiB+zj7YsKrIWkpH2xjLGyTDbadIhAhBXNmMDI2J68rgNw==
	MEUCIQDk7H3OWOrGtYpr/CzoOHb8XE61J9MqVfv89eskQWvTLAIgVaPsQcfA8BRUEw4ZSu17jpEn06fVh5lEZ+VtfzjwG+c=
	MEUCIB112D72ACL1qdv/QUdpC0em3EcT+FqkyNq4k7nHoNePAiEA2C88hUuI8sfN7dbJR3+FTJ8E1EXrQ2NEUP213orA4fI=
	MEUCIByYSRXK/U/r6t7FY1nosZbhNJSrMXuK8+Ra92DtpySGAiEAwHMuh3V1DDAthWcF8DjLQO7fNPRyZciUolESxkHLqJo=
	MEYCIQDKFqxbhla/4TQ/Y7st0aJAF8fXLdWulgsk3cA1F2YNMgIhAKMmGToo2il3Pljs8q2QeNksGQOjjX17smzEucVLa4HM
	MEQCIHRUQoS2UnE1DJcmh9X7QAniNJqHziiFDXbsSKw4C4PUAiBso6WXhRCE+0hWjgBqzt8A6xf8iIxq2JfslDGG+jhTCA==
	MEUCIQCGEavNNxJzrMzj5bbHAXjqOhJzFod65n3QAPKQ0HPa8AIgAOk33lRRDct8YzV8FAohInc6Tzzp4xwyu9YPUdAzW5Y=
	MEUCIGxzeZIXC+Q1FTm8nxvHszPa5rbNMLMOMA+Z9ZQuzYfVAiEA7yxf6rvUu5lpSIPDZUO/aKjZLodjNA6X7cRduvoqSFM=
	MEUCIQDtKHs4g5IGdZA4VMlNgtAPiuGponug19XbC2ku5mSlJwIgFHpDO+ebk1zaVJCadKGU2WxwqyIiczXJwrijWeDOqzM=
	MEUCIQDzwDAJUCa163Uja7HIv3QohvhfGqrVazuixcGK33gW2QIgXjNlEj4LOtXGKZMNTIbAeiZbNA9/Ch0GjkmrIN2lYl8=
	MEYCIQCyTDeJ2KuNcCSwWQ1j+iRwTr68M5Gp4BiYS6v4kpzgAQIhALRjXepHBmxGBgj4ymAeMF1tgXfp8RHxFJwB9qHU5TAh
	MEUCIDRmO/vNY0jmLWC0dV0K1Qixlos0xh/mOMLLb473ATCUAiEA922TyIB5LioNzL4ZLOg+FzulC5SGp4F7B1meEpg6S6M=
	MEYCIQCVEJPF7gSuCm/bYO7vgG8+zLgAuIOZ+dqYqYK/Au1K4AIhALqeQFbHkQl061mwodkUAmdn/2hMXiHKY2u2X+dTX62G
	MEUCIErFV8SMwTC0128KZrP5FCzRGG0lXqVnDNrGEBuBYeUbAiEA6UAWmj9nlx+MQKTYYn378Zelym6wish8DFLpHaS/PJQ=
	MEUCIQD5YTJ860vZAZFNtzY8A9epR+XsBPEi2Sf/ngMBkyHxZAIgedqFRFSLZXxgNq7kVkhBedvL8qGeS4HUsSM55Fc2+qk=
	MEUCIHs50gGTHWgQ+g1KadGZ6zyXg3QQDauULQm+UbC+2xSiAiEAqRTQulEnetzaUssFYTtF/aZRAwYZwlUAy6eTqCel+rU=
	MEUCIQC6AjgtBj64eQuDc5cXhaTfHoPVqzai2jiRpTgLP/LvpgIgDIc2+eUfGQ9qX9yCpSGfIQ3Gug67j2gDZ69+/pJRdHk=
	MEUCIQC9px90GlTREBN8K/paxII9WI8wTc1xcT1hre7WitvFWAIge+F05Nz0R/DfiaoaUG3tcyd5uECkStXnyWCrrfX3Fps=
	MEUCIDVLP+WDpTcOu2Oxndy65BPDZ0Nkt1jyk1y5OfjylXzPAiEArpbpzQmnTK6PjGA42a5d8AyHvn3tdljFYtLMNMAAz1U=
	MEUCIG7rXcKiBSBG4P/Vj7Bfh5yYzFtso7lgAtZuHF4gjdwZAiEA7zKifX2KFZRUpn8nH/AFCmBQAjfhGG0No4eF+DMmMiI=
	MEYCIQCii44AhjfkTCl+v2udARDE0c0ssnSmz4sAzT9Rq/0IqgIhAKBfiaMO+S3c1sK8PvIGwZDrv2usQfVU1YJHCBFXRJ/A
	MEQCIHVqIeu+yYlm2HHJ6xOnz6OI0GmMmEowimzXHPoPWRh+AiBF/a2z/a3ND0jZCph5hdw3cQFbTeeWRo4cUa2GPSylCQ==
	MEQCIFswuS+a0mM2dsPtfRVjDkPKld2Hu6YYVWiZ4Io5LmCWAiAiLg+6ElEAAdZ+CsJZyOGsu1F9TXvyPH8/NzgRT13A4g==
	MEQCIELbzzUk/zd/lv5/DOd3Tg1m8Y5mBhmGDaFRBXg2rEbHAiBrxGAprooPB5iRqZ+Ia/YYpEfc0VivOETyj0EIlW/5Lw==`
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, PK, hash[:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	sigHash := base64.StdEncoding.EncodeToString(sig)
	w.Write([]byte(sigHash))
}