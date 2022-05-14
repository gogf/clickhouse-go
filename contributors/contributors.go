// Licensed to ClickHouse, Inc. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. ClickHouse, Inc. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package contributors

import (
	//_ "embed"
	"strings"
)

var source = `Abraham Adberstein <abraham@cloudflare.com>
Aleksandr Petrukhin <a.petrukhin@city-mobil.ru>
Alex Bocharov <ab@cloudflare.com>
Alex Yang <ayyang9@gmail.com>
Alexander Chumakov <bestak47@gmail.com>
Alexander Obukhov <dev@sprql.space>
Alexey Milovidov <milovidov@clickhouse.com>
Alexey Palazhchenko <alexey.palazhchenko@gmail.com>
Alvaro Tuso <alvarotuso@hotmail.com>
Andrey Ustinov <ustinov.post@gmail.com>
Ashish Gaurav <ashishgaurav.iitd@gmail.com>
Benjamin Rupp <brupp@ciena.com>
Cem Sancak <cem.sancak.90@gmail.com>
Chao Wang <chaowang@uber.com>
Chris Duncan <veqryn@hotmail.com>
Dale McDiarmid <dale@clickhouse.com>
Damir Sayfutdinov <sayfutdinov@selectel.ru>
Dan Walters <dan@walters.io>
Daniel Bershatsky <daniel.bershatsky@skolkovotech.ru>
Danila Migalin <miga@uber.com>
Danny.Dunn <danny@DannyDunndeMBP.lan>
Darío <dgrripoll@gmail.com>
Denis Krivak <dokrivak@avito.ru>
Derek Perkins <derek@derekperkins.com>
Dmitry Markov <dmitri__13@mail.ru>
Dmitry Ponomarev <demdxx@gmail.com>
Dmitry Ponomarev <demdxx@trafficstars.com>
Egor.Gorlin <Egor.Gorlin@forextime.com>
Eugene Formanenko <mo4islona@gmail.com>
Evan Au <au.liangjun@gmail.com>
Ewan <ewan.p.walker@gmail.com>
Florian Lehner <flehner@optimyze.cloud>
Félix Mattrat <felix@messagebird.com>
Geoff Genz <geoff@clickhouse.com>
Ian McGraw <ian@arthur.ai>
Ivan Blinkov <github@blinkov.ru>
Ivan Ivanov <ivanov@corp.mail.ru>
Jake Sylvestre <jakesyl@gmail.com>
Jakub Chábek <jakub.chabek@cdn77.com>
James Hartig <fastest963@gmail.com>
Jan Was <j.was@f5.com>
Jeehoon Kim <jeehooni@gmail.com>
John Troy <john@noxi.us>
Jon Aquino <jonathan.aquino@adroll.com>
LI Tao <litao.91@bytedance.com>
Maksim Sokolnikov <stuffsweep@gmail.com>
Marek Vavruša <marek@vavrusa.com>
Marek Vavruša <mvavrusa@cloudflare.com>
Mark Andrus Roberts <markandrusroberts@gmail.com>
Michael Vigovsky <upliner@gmail.com>
Nay Linn <nlinn@arista.com>
Oleg Strokachuk <oleg.strokachuk@vizor-games.com>
Richard Artoul <richardartoul@gmail.com>
Robert Sköld <robert@department.se>
Robin Hahling <robin.hahling@gw-computing.net>
Ross Rothenstine <rossr@unity3d.com>
Sergei Sobolev <ssobolev@ozon.ru>
Sergey Melekhin <sergey@melekhin.me>
Taras Matsyk <taras@mm.st>
Thibault Deutsch <thibault@arista.com>
Tsimafei Bredau <t.bredau@affise.com>
Vespertinus <rrvespertinus@gmail.com>
Yury Korolev <yury.king@gmail.com>
Yury Yurochko <y.yurochko@iconic.vc>
alex <alex@localhost.localdomain>
anton troyanov <anton@troyanov.net>
chengzhi <chengzhi@shinnytech.com>
dmitry kuzmin <dmitry.kuzmin@exness.com>
gaetan.rizio <gaetan.rizio@contentsquare-ext.com>
hulb <hulb@live.cn>
ilker moral <ilker.moral@comodo.com>
jiyongwang <jiyongwang@freewheel.tv>
kshvakov <shvakov@gmail.com>
neverlee <neverlea@foxmail.com>
nevseliev <nevseliev@life-team.net>
pavel raskin <f0ma@inbox.ru>
sundy-li <543950155@qq.com>
vahid sohrabloo <vahid4134@gmail.com>
vasily.popov <vasily.popov@arrival.com>
vl4deee11 <boi4enkovlad@yandex>
vl4deee11 <vl4deee11@gmail.com>
vladislav doster <mvdoster@gmail.com>
vvoronin <voronin@x12.su>
yuankun <rogeryk@outlook.com>
zxc111 <zxc9007@gmail.com>
李盼 <lipan@sunteng.com>`

var List = strings.Split(source, "\n")
