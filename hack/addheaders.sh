#!/bin/bash
#     Copyright 2020 Aicura Nexus Client and/or its authors
#
#     This file is part of Aicura Nexus Client.
#
#     Aicura Nexus Client is free software: you can redistribute it and/or modify
#     it under the terms of the GNU Lesser General Public License as published by
#     the Free Software Foundation, either version 3 of the License, or
#     (at your option) any later version.
#
#     Aicura Nexus Client is distributed in the hope that it will be useful,
#     but WITHOUT ANY WARRANTY; without even the implied warranty of
#     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#     GNU Lesser General Public License for more details.
#
#     You should have received a copy of the GNU Lesser General Public License
#     along with Aicura Nexus Client.  If not, see <https://www.gnu.org/licenses/>.

which ./bin/addlicense >/dev/null || go build -o ./bin/addlicense github.com/google/addlicense

# https://github.com/google/addlicense
# https://www.gnu.org/licenses/gpl-howto.en.html

./bin/addlicense -c "Aicura Nexus Client and/or its authors" -f LICENSE_NOTICE hack nexus
